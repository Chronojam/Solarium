package interfaces

import "time"

// PlanetaryBody is the interface which a solar system can accept.
type PlanetaryBody interface {
	// Returns the map representing the properties of this body.
	Properties() Properties
	Name() string
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

	Description() string
}

type Properties interface {
	GetTemperature() int
	SetTemperature(v int) int

	GetRadiation() int
	SetRadiation(v int) int

	GetAtmosphericDensity() int
	SetAtmosphericDensity(v int)

	GetDistanceFromSun() int
	SetDistanceFromSun(v int)
}
