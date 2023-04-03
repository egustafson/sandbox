package protogw

import (
	"encoding/json"
	"fmt"

	"github.com/egustafson/sandbox/go/kv-gateway/kv"
	"github.com/google/uuid"
)

/* This package will demonstrate the Gateway pattern and model a small number of
 * made up domain objects that can be persisted and reconstituted through the
 * Gateway.
 */

type ProtoGateway interface {
	Close()
	SetProducer(p Producer) error
	GetProducer(id uuid.UUID) (Producer, error)
	GetProducers() ([]Producer, error)
	SendLogRecord(lr LogRecord) error
	GetAllLogs() ([]LogRecord, error)
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

func (pg *kvProtoGateway) SetProducer(p Producer) error {
	return pg.setObject(&p)
}

func (pg *kvProtoGateway) GetProducer(id uuid.UUID) (p Producer, err error) {
	v, err := pg.getObject((&p).MkKey(id.String()))
	if err != nil {
		return
	}
	err = json.Unmarshal(v, &p)
	return
}

func (pg *kvProtoGateway) GetProducers() ([]Producer, error) {
	objs, err := pg.getObjSlice((&Producer{}).GwPrefix())
	if err != nil {
		return nil, err
	}
	producers := make([]Producer, len(objs))
	for i, o := range objs {
		if err = json.Unmarshal(o, &(producers[i])); err != nil {
			return nil, err
		}
	}
	return producers, nil
}

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
