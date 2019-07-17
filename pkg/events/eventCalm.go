package events

import (
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// Calm, nothing is happening right now.
type Calm struct {
	expiresAt  time.Time
	TimeLength time.Duration
	Desc       string
}

func (c *Calm) Apply(p []interfaces.PlanetaryBody) {}
func (c *Calm) Cease(p []interfaces.PlanetaryBody) {}
func (c *Calm) Duration() time.Duration {
	return c.TimeLength
}
func (c *Calm) NextEvent() interfaces.SystemEvent {
	return RandomEvent()
}
func (c *Calm) Description() string {
	if c.Desc != "" {
		return c.Desc
	}
	return "all is calm"
}
