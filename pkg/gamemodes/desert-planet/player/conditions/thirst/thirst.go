package thirst

import "github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player/conditions"

const (
	defaultThirstGained = 2
)

func New() *Thirst {
	return &Thirst{
		v:          100,
		multiplier: 1,
	}
}

type Thirst struct {
	v          int
	multiplier int
}

func (t *Thirst) State() conditions.State {
	if t.v > 80 {
		return conditions.Good
	}
	if t.v <= 80 && t.v >= 50 {
		return conditions.Normal
	}
	if t.v < 50 && t.v >= 20 {
		return conditions.NeedsAttention
	}
	if t.v < 20 {
		return conditions.Fatal
	}
	return conditions.Normal
}

func (t *Thirst) Simulate() {
	t.v = t.v - defaultThirstGained*t.multiplier
}
func (t *Thirst) SetMultiplier(i int) {
	t.multiplier = i
}
func (t *Thirst) SetValue(i int) {
	t.v = i
}
func (t *Thirst) FatalText() string {
	return `
Died of thirst
	`
}
