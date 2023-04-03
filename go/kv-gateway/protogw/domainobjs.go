package protogw

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Producer struct {
	ID    uuid.UUID `yaml:"id" json:"id"`
	Name  string    `yaml:"name" json:"name"`
	Value int       `yaml:"val" json:"val"`
}

// static check: a *Producer isA GatewayObj
var _ GatewayObj = (*Producer)(nil)

type LogRecord struct {
	TimeStamp time.Time `yaml:"ts" json:"ts"`
	Producer  uuid.UUID `yaml:"produced-by" json:"produced-by"`
	Message   string    `yaml:"message" json:"message"`
}

// static check: a *LogRecord isA GatewayObj
var _ GatewayObj = (*LogRecord)(nil)

// --  Producer Implementation  -------------------------------------

func (p *Producer) GwPrefix() string       { return "producer" }
func (p *Producer) MkKey(id string) string { return fmt.Sprintf("%s/%s", p.GwPrefix(), id) }
func (p *Producer) GwKey() string          { return p.MkKey(p.ID.String()) }

// --  LogRecord Implementation  ------------------------------------

func (lr *LogRecord) GwPrefix() string { return "log" }

func (lr *LogRecord) GwKey() string {
	// log/<timestamp-rfc3339nano>/<producer-id>
	return fmt.Sprintf("%s/%s/%s",
		lr.GwPrefix(),
		lr.TimeStamp.UTC().Format(time.RFC3339Nano),
		lr.Producer.String())
}
