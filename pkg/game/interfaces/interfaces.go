package interfaces

// PlanetaryBody is the interface which a solar system can accept.
type PlanetaryBody interface {
	// Returns the map representing the properties of this body.
	Name() string
	AddEffect(p PlanetaryEffect)
	RemoveEffect(p PlanetaryEffect)
}

type GameEvent struct {
	Description string
}

type PlanetaryEffect interface {
	ID() string
	Name() string
	Desc() string
}
