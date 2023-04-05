package protogw_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/egustafson/sandbox/go/kv-gateway/protogw"
)

func TestComponent_GwKey(t *testing.T) {
	c := &protogw.Component{
		ID: uuid.New(),
	}

	expectedKey := fmt.Sprintf("%s/%s", c.GwPrefix(), c.ID.String())
	assert.Equal(t, expectedKey, c.GwKey())
}

func TestComponent_MkKey(t *testing.T) {
	c := &protogw.Component{}
	id := uuid.New().String()

	expectedKey := fmt.Sprintf("%s/%s", c.GwPrefix(), id)
	assert.Equal(t, expectedKey, c.MkKey(id))
}

func TestLogRecord_GwKey(t *testing.T) {
	lr := &protogw.LogRecord{
		TimeStamp: time.Now(),
		Sender:    uuid.New(),
	}

	expectedKey := fmt.Sprintf("%s/%s/%s",
		lr.GwPrefix(),
		lr.TimeStamp.UTC().Format(time.RFC3339Nano),
		lr.Sender.String())
	assert.Equal(t, expectedKey, lr.GwKey())
}

func TestGaugeState_GwKey(t *testing.T) {
	g := &protogw.GaugeState{
		OwnerID:     uuid.New(),
		Description: "test description",
		Value:       21,
	}

	expectedKey := fmt.Sprintf("%s/%s", g.GwPrefix(), g.OwnerID.String())
	assert.Equal(t, expectedKey, g.GwKey())
}
