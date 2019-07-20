package player

import (
	"github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/player/conditions"
)

type condition interface {
	State() conditions.State
	Simulate()
	SetMultiplier(i int)
	FatalText() string
	SetValue(i int)
}

type job interface {
	Simulate()
}

// Player represents an human player.
type Player struct {
	ID   string
	Name string

	Conditions []condition
	Job        job
}
