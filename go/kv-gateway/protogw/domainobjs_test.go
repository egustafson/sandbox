package protogw_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/egustafson/sandbox/go/kv-gateway/protogw"
)

func TestProducer_GwKey(t *testing.T) {
	p := &protogw.Producer{
		ID: uuid.New(),
	}

	expectedKey := fmt.Sprintf("%s/%s", p.GwPrefix(), p.ID.String())
	assert.Equal(t, expectedKey, p.GwKey())
}

func TestProducer_MkKey(t *testing.T) {
	p := &protogw.Producer{}
	id := uuid.New().String()

	expectedKey := fmt.Sprintf("%s/%s", p.GwPrefix(), id)
	assert.Equal(t, expectedKey, p.MkKey(id))
}

func TestLogRecord_GwKey(t *testing.T) {
	lr := &protogw.LogRecord{
		TimeStamp: time.Now(),
		Producer:  uuid.New(),
	}

	expectedKey := fmt.Sprintf("%s/%s/%s",
		lr.GwPrefix(),
		lr.TimeStamp.UTC().Format(time.RFC3339Nano),
		lr.Producer.String())
	assert.Equal(t, expectedKey, lr.GwKey())
}
