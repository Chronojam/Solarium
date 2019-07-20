package food

import (
	"math/rand"
)

type Food struct {
	RemainingTime int
	AmountGained  int
}

func New(useFuel bool) *Food {
	DistanceTillFood := rand.Intn(3)
	if useFuel {
		DistanceTillFood = DistanceTillFood / 2
	}
	return &Food{
		RemainingTime: DistanceTillFood,
		AmountGained:  rand.Intn(3),
	}
}

func (w *Food) Simulate() {
	w.RemainingTime--
}

func (w *Food) Done() bool {
	return w.RemainingTime == 0
}
func (w *Food) Amount() int {
	return w.AmountGained
}
