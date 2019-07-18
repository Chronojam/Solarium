package interfaces

import (
	"time"
)

// PlanetaryBody is the interface which a solar system can accept.
type PlanetaryBody interface {
	// Returns the map representing the properties of this body.
	Name() string
	AddEffect(p PlanetaryEffect)
	RemoveEffect(p PlanetaryEffect)
}

type SystemEvent interface {
	// Applies the event to a given planetaryBody
	Apply(p []PlanetaryBody)
	// Ceases the event on a given planetaryBody
	Cease(p []PlanetaryBody)
	// ExpiresAt returns the time that this event expires at.
	Duration() time.Duration
	// NextEvent returns the next logical event to this one
	NextEvent() SystemEvent
	// Desc
	Desc() string
}

type PlanetaryEffect interface {
	ID() string
	Name() string
	Desc() string
}
