package protogw

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/egustafson/werks/kv"
	"github.com/google/uuid"
)

/* This package will demonstrate the Gateway pattern and model a small number of
 * made up domain objects that can be persisted and reconstituted through the
 * Gateway.
 */

type ProtoGateway interface {
	Close()

	SetComponent(c Component) error
	GetComponent(id uuid.UUID) (Component, error)
	GetAllComponents() ([]Component, error)

	CreateGauge(id uuid.UUID, desc string) error
	SetGauge(id uuid.UUID, val int) error
	ObserveGauge(ctx context.Context, id uuid.UUID) (<-chan GaugeState, error)

	SendLogRecord(lr LogRecord) error
	GetAllLogs() ([]LogRecord, error)
	ObserveLogs(ctx context.Context) (<-chan LogRecord, error)
}

type GatewayObj interface {
	GwPrefix() string // return the gateway scoped prefix of the object class
	GwKey() string    // return the storage key for the object
}

type kvProtoGateway struct {
	kvstore   kv.KV
	keyPrefix string
}

// static check: a *kvProtoGateway isA ProtoGateway
var _ ProtoGateway = (*kvProtoGateway)(nil)

func MakeProtoGateway(gwPrefix string, kvstore kv.KV) ProtoGateway {
	pg := &kvProtoGateway{
		kvstore:   kvstore,
		keyPrefix: gwPrefix,
	}
	return pg
}

func (pg *kvProtoGateway) Close() {
	pg.kvstore.Close()
}

func (pg *kvProtoGateway) setObject(obj GatewayObj) error {
	v, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	k := kv.Key(fmt.Sprintf("%s/%s", pg.keyPrefix, obj.GwKey()))
	if err = pg.kvstore.Put(k, v); err != nil {
		return err
	}
	return err
}

func (pg *kvProtoGateway) getObject(key string) (kv.Value, error) {
	k := fmt.Sprintf("%s/%s", pg.keyPrefix, key)
	if v, err := pg.kvstore.Get(kv.Key(k)); err != nil {
		return nil, err
	} else {
		return v, nil
	}
}

func (pg *kvProtoGateway) getObjSlice(keyPrefix string) ([]kv.Value, error) {
	k := fmt.Sprintf("%s/%s", pg.keyPrefix, keyPrefix)
	kvs, err := pg.kvstore.GetPrefix(kv.Key(k))
	if err != nil {
		return nil, err
	}
	vals := make([]kv.Value, 0, len(kvs))
	for _, entry := range kvs {
		vals = append(vals, entry.V)
	}
	return vals, nil
}

// --  Component methods  ------------------------------------------------

func (pg *kvProtoGateway) SetComponent(c Component) error {
	return pg.setObject(&c)
}

func (pg *kvProtoGateway) GetComponent(id uuid.UUID) (c Component, err error) {
	v, err := pg.getObject((&c).MkKey(id.String()))
	if err != nil {
		return
	}
	err = json.Unmarshal(v, &c)
	return
}

func (pg *kvProtoGateway) GetAllComponents() ([]Component, error) {
	objs, err := pg.getObjSlice((&Component{}).GwPrefix())
	if err != nil {
		return nil, err
	}
	components := make([]Component, len(objs))
	for i, o := range objs {
		if err = json.Unmarshal(o, &(components[i])); err != nil {
			return nil, err
		}
	}
	return components, nil
}

// --  LogRecord methods  ------------------------------------------------

func (pg *kvProtoGateway) SendLogRecord(lr LogRecord) error {
	return pg.setObject(&lr)
}

func (pg *kvProtoGateway) GetAllLogs() ([]LogRecord, error) {
	objs, err := pg.getObjSlice((&LogRecord{}).GwPrefix())
	if err != nil {
		return nil, err
	}
	recs := make([]LogRecord, len(objs))
	for i, o := range objs {
		if err = json.Unmarshal(o, &(recs[i])); err != nil {
			return nil, err
		}
	}
	return recs, nil
}

func (pg *kvProtoGateway) ObserveLogs(ctx context.Context) (<-chan LogRecord, error) {
	var err error
	var evCh <-chan []kv.Event
	prefix := (&LogRecord{}).GwPrefix()
	if evCh, err = pg.kvstore.WatchPrefix(ctx, kv.Key(prefix)); err != nil {
		return nil, err
	}
	logsCh := make(chan LogRecord, 10)
	go func() {
		defer close(logsCh)
		for {
			select {
			case events := <-evCh:
				var lr LogRecord
				for _, ev := range events {
					switch ev.EventType {
					case kv.PutEvent:
						if err := json.Unmarshal(ev.Kv.V, &lr); err == nil {
							logsCh <- lr
						}
					default:
						// should never happen - log or something
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return logsCh, nil
}

// --  GaugeState methods  -----------------------------------------------

func (pg *kvProtoGateway) CreateGauge(id uuid.UUID, desc string) error {
	g := &GaugeState{
		OwnerID:     id,
		Description: desc,
	}
	return pg.setObject(g)
}

func (pg *kvProtoGateway) getGauge(id uuid.UUID) (g *GaugeState, err error) {
	v, err := pg.getObject(g.MkKey(id.String()))
	if err != nil {
		return
	}
	g = new(GaugeState)
	err = json.Unmarshal(v, g)
	return
}

func (pg *kvProtoGateway) SetGauge(id uuid.UUID, val int) error {
	g, err := pg.getGauge(id)
	if err != nil {
		return err
	}
	g.Value = val
	return pg.setObject(g)
}

func (pg *kvProtoGateway) ObserveGauge(ctx context.Context, id uuid.UUID) (<-chan GaugeState, error) {
	var err error
	if _, err = pg.getGauge(id); err != nil {
		return nil, err
	}
	key := (&GaugeState{}).MkKey(id.String())
	var evCh <-chan []kv.Event
	if evCh, err = pg.kvstore.Watch(ctx, kv.Key(key)); err != nil {
		return nil, err
	}
	gaugeCh := make(chan GaugeState, 10)
	go func() {
		defer close(gaugeCh)
		for {
			select {
			case events := <-evCh:
				for _, ev := range events {
					var g GaugeState
					switch ev.EventType {
					case kv.PutEvent:
						if err := json.Unmarshal(ev.Kv.V, &g); err == nil {
							gaugeCh <- g
						}
					case kv.DelEvent:
						if err := json.Unmarshal(ev.PrevKv.V, &g); err == nil {
							g.Value = 0
							gaugeCh <- g
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return gaugeCh, nil
}
