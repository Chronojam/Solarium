package water

import (
	"math/rand"
)

type Water struct {
	RemainingTime int
	AmountGained  int
}

func New(useFuel bool) *Water {
	DistanceTillWater := rand.Intn(3)
	if useFuel {
		DistanceTillWater = DistanceTillWater / 2
	}
	return &Water{
		RemainingTime: DistanceTillWater,
		AmountGained:  rand.Intn(3) + 1,
	}
}

func (w *Water) Simulate() {
	w.RemainingTime--
}

func (w *Water) Done() bool {
	return w.RemainingTime == 0
}
func (w *Water) Amount() int {
	return w.AmountGained
}
