package events

import (
	"log"
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// EquilibriumConstant will try to restore planets to their natural state
// over time
type EquilibriumConstant struct{}

func (c *EquilibriumConstant) Apply(p []interfaces.PlanetaryBody) {
	for _, body := range p {
		prop := body.Properties()
		log.Printf("Setting temp on: %v", body.Name())
		if prop.GetTemperature() > prop.GetDistanceFromSun()*prop.GetAtmosphericDensity()/100 {
			log.Printf("-T: %v", prop.GetTemperature()-5)
			prop.SetTemperature(prop.GetTemperature() - 5)
		} else {
			log.Printf("+T: %v", prop.GetTemperature()+5)
			prop.SetTemperature(prop.GetTemperature() + 5)
		}
		if prop.GetRadiation() > prop.GetDistanceFromSun()*2 {
			prop.SetRadiation(prop.GetRadiation() - 5)
		} else {
			prop.SetRadiation(prop.GetRadiation() + 5)
		}
	}
}
func (c *EquilibriumConstant) Cease(p []interfaces.PlanetaryBody) {}
func (c *EquilibriumConstant) Duration() time.Duration {
	if DEBUG_MODE {
		return time.Millisecond * 500
	}
	return time.Second * 10
}
func (c *EquilibriumConstant) NextEvent() interfaces.SystemEvent {
	return &EquilibriumConstant{}
}
func (c *EquilibriumConstant) Description() string {
	return "the universe equalizes"
}
