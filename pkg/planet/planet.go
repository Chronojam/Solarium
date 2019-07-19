package planet

import (
	"math/rand"

	"github.com/chronojam/solarium/pkg/planet/effects"
	"github.com/chronojam/solarium/pkg/planet/hardware"
	"github.com/chronojam/solarium/pkg/planet/namegenerator"
	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// planet represents the world on which a given player controls.
type planet struct {
	name, owner string
	effects     []interfaces.PlanetaryEffect

	Temperature   int
	Radioactivity int

	DistanceFromSun    int
	AtmosphericDensity int
}

// New returns a new planet with the given properties
func New(name string) *planet {
	return &planet{
		name: name,
	}
}

// NewRandom returns a new planet with a random set
// of properties.
func NewRandom(name, owner string) *planet {
	if name == "" {
		name = namegenerator.GenerateNew()
	}
	MinDistance := 1
	MaxDistance := 50 - MinDistance

	MaxAtmosphere := 300

	p := &planet{
		name:               name,
		DistanceFromSun:    MinDistance + rand.Intn(MaxDistance),
		AtmosphericDensity: rand.Intn(MaxAtmosphere),
	}
	p.UpdateRadioactivity()
	p.UpdateTemperature()

	return p
}

func (p *planet) AddEffect(e interfaces.PlanetaryEffect) {
	p.effects = append(p.effects, e)
}
func (p *planet) RemoveEffect(e interfaces.PlanetaryEffect) {
	indexToRemove := -1
	for index, ef := range p.effects {
		if ef.ID() == e.ID() {
			indexToRemove = index
		}
	}
	if indexToRemove == -1 {
		// Doesnt exist in the index.
		return
	}
	p.effects = append(p.effects[:indexToRemove], p.effects[indexToRemove+1:]...)
}

func (p *planet) UpdateTemperature() {
	defaultTemp := (p.AtmosphericDensity ^ 2) - p.DistanceFromSun
	proposedAlteration := 0
	for _, e := range p.effects {
		if val, ok := e.(effects.TemperatureEffect); ok {
			// Allow each temp effect to do something with the
			proposedAlteration += val.AlterTemp(defaultTemp)
		}
	}
	// After all the effects have taken place, we can allow hardware
	// to amplify or mitigate the effects
	for _, e := range p.effects {
		if val, ok := e.(hardware.TemperatureEffect); ok {
			proposedAlteration = val.HWAlterTemp(proposedAlteration)
		}
	}
	p.Temperature = defaultTemp + proposedAlteration
}

func (p *planet) UpdateRadioactivity() {
	defaultRads := p.DistanceFromSun / 2
	proposedAlteration := 0
	for _, e := range p.effects {
		if val, ok := e.(effects.RadioactivityEffect); ok {
			// Allow each temp effect to do something with the
			proposedAlteration += val.AlterRads(defaultRads)
		}
	}
	// After all the effects have taken place, we can allow hardware
	// to amplify or mitigate the effects
	for _, e := range p.effects {
		if val, ok := e.(hardware.RadioactivityEffect); ok {
			proposedAlteration = val.HWAlterRads(proposedAlteration)
		}
	}

	p.Radioactivity = defaultRads + proposedAlteration
}

func (p *planet) Name() string {
	return p.name
}
