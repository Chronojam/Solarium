package main

import (
	"github.com/chronojam/solarium/pkg/planet"
	"github.com/chronojam/solarium/pkg/system"
)

func main() {
	s := system.New()
	s.AddBody(
		planet.NewRandom("The Quickening Of Doom"),
	)
	s.Simulate()

}
