package protogw

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// --  Component  --------------------------------------------------------

// Component is a generic object
type Component struct {
	ID   uuid.UUID `yaml:"id" json:"id"`
	Name string    `yaml:"name" json:"name"`
}

// static check: a *Producer isA GatewayObj
var _ GatewayObj = (*Component)(nil)

func (c *Component) GwPrefix() string       { return "producer" }
func (c *Component) MkKey(id string) string { return fmt.Sprintf("%s/%s", c.GwPrefix(), id) }
func (c *Component) GwKey() string          { return c.MkKey(c.ID.String()) }

// --  LogRecord  --------------------------------------------------------

// LogRecord is a single log record
type LogRecord struct {
	TimeStamp time.Time `yaml:"ts" json:"ts"`
	Sender    uuid.UUID `yaml:"sender" json:"sender"`
	Message   string    `yaml:"message" json:"message"`
}

// static check: a *LogRecord isA GatewayObj
var _ GatewayObj = (*LogRecord)(nil)

func (lr *LogRecord) GwPrefix() string { return "log" }

func (lr *LogRecord) GwKey() string {
	// log/<timestamp-rfc3339nano>/<producer-id>
	return fmt.Sprintf("%s/%s/%s",
		lr.GwPrefix(),
		lr.TimeStamp.UTC().Format(time.RFC3339Nano),
		lr.Sender.String())
}

// --  GaugeState  -------------------------------------------------------

// GaugeState represents the value of some general measurement.
type GaugeState struct {
	OwnerID     uuid.UUID `yaml:"owner-id" json:"owner-id"`
	Description string    `yaml:"description" json:"description"`
	Value       int       `yaml:"value" json:"value"` // the gauge's "value"
}

// static check: a *GaugeState isA GatewayObj
var _ GatewayObj = (*GaugeState)(nil)

func (g *GaugeState) GwPrefix() string       { return "gauge" }
func (g *GaugeState) MkKey(id string) string { return fmt.Sprintf("%s/%s", g.GwPrefix(), id) }
func (g *GaugeState) GwKey() string          { return g.MkKey(g.OwnerID.String()) }
