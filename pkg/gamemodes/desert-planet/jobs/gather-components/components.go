package components

import (
	"math/rand"
)

type Components struct {
	RemainingTime int
	AmountGained  int
}

func New(useFuel bool) *Components {
	DistanceTillComponent := rand.Intn(3)
	if useFuel {
		DistanceTillComponent = DistanceTillComponent / 2
	}
	return &Components{
		RemainingTime: DistanceTillComponent,
		AmountGained:  1,
	}
}

func (w *Components) Simulate() {
	w.RemainingTime--
}

func (w *Components) Done() bool {
	return w.RemainingTime == 0
}
func (w *Components) Amount() int {
	return w.AmountGained
}
