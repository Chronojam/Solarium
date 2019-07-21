package hunger

import (
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player/conditions"
)

const defaultHungerGained = 1

type Hunger struct {
	v          int
	multiplier int
}

func New() *Hunger {
	return &Hunger{
		v:          100,
		multiplier: 1,
	}
}

func (t *Hunger) State() conditions.State {
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

func (t *Hunger) Simulate() {
	t.v = t.v - defaultHungerGained*t.multiplier
}
func (t *Hunger) SetMultiplier(i int) {
	t.multiplier = i
}
func (t *Hunger) SetValue(i int) {
	t.v = i
}
func (t *Hunger) FatalText() string {
	return `
Died of hunger
	`
}
