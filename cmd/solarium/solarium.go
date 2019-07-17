package main

import (
	"github.com/chronojam/solarium/pkg/planet"
	"github.com/chronojam/solarium/pkg/planet/hardware"
	"github.com/chronojam/solarium/pkg/system"
)

func main() {
	s := system.New()
	s.AddBody(
		planet.New(
			"The Quickening Of Doom",
			&planet.Properties{
				DistanceFromSun:   100,
				AtmosphereDensity: 100,
				Hardware: []planet.PlanetaryHardware{
					&hardware.HeatShields{Quality: 0.25},
				},
			},
		),
	)
	s.Simulate()

}
