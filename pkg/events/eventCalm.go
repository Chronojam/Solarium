package events

import (
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// Calm, nothing is happening right now.
type Calm struct {
	expiresAt   time.Time
	TimeLength  time.Duration
	Description string
}

func (c *Calm) Apply(p []interfaces.PlanetaryBody) {}
func (c *Calm) Cease(p []interfaces.PlanetaryBody) {}
func (c *Calm) Duration() time.Duration {
	return c.TimeLength
}
func (c *Calm) NextEvent() interfaces.SystemEvent {
	return RandomEvent()
}
func (c *Calm) Desc() string {
	if c.Description != "" {
		return c.Description
	}
	return "All is calm."
}
