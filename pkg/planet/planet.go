package planet

import (
	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// planet represents the world on which a given player controls.
type planet struct {
	properties *Properties
	name       string
}

// New returns a new planet with the given properties
func New(name string, prop *Properties) *planet {
	prop.Temperature = prop.DistanceFromSun * prop.AtmosphereDensity / 100
	prop.Radiation = prop.DistanceFromSun * 2

	return &planet{
		name:       name,
		properties: prop,
	}
}

// NewRandom returns a new planet with a random set
// of properties.
func NewRandom() *planet {
	return &planet{}
}

func (p *planet) Properties() interfaces.Properties {
	return p.properties
}

func (p *planet) Name() string {
	return p.name
}
