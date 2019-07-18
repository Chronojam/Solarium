package properties

import "log"

// Properties are the properties of a planet.
type Properties struct {
	DistanceFromSun   int
	AtmosphereDensity int

	Temperature int
	Radiation   int

	Hardware []PlanetaryHardware
}

func (p *Properties) GetTemperature() int {
	return p.Temperature
}
func (p *Properties) SetTemperature(v int) int {
	// How much are we proposing to change the temp by?
	change := v - p.Temperature
	log.Printf("Temp Change: %v\n", change)
	for _, h := range p.Hardware {
		if val, ok := h.(TemperatureHardware); ok {
			// Allow each piece of temp hardward to alter the effects
			// of the temperature change

			// this needs the delta of change
			change = val.AlterTemp(change)
		}
	}
	log.Printf("Actual Temp: %v\n", p.Temperature+change)
	p.Temperature = p.Temperature + change
	return change
}

func (p *Properties) GetRadiation() int {
	return p.Radiation
}
func (p *Properties) SetRadiation(v int) int {
	change := v - p.Radiation
	for _, h := range p.Hardware {
		if val, ok := h.(RadiationHardware); ok {
			change = val.AlterRads(change)
		}
	}
	p.Temperature = p.Radiation + change
	return change
}

func (p *Properties) GetAtmosphericDensity() int {
	return p.DistanceFromSun
}
func (p *Properties) SetAtmosphericDensity(v int) {
	p.AtmosphereDensity = v
}

func (p *Properties) GetDistanceFromSun() int {
	return p.DistanceFromSun
}
func (p *Properties) SetDistanceFromSun(v int) {
	p.DistanceFromSun = v
}

type PlanetaryHardware interface{}
type TemperatureHardware interface {
	AlterTemp(v int) int
}
type RadiationHardware interface {
	AlterRads(v int) int
}
