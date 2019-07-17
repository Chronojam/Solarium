package events

import (
	"math/rand"
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// SolarStorm brews up a horrid storm, increasing tempreture and radiation
// on the body
type SolarStorm struct {
	TimeLength time.Duration
	Intensity  int
	Desc       string
}

func (s *SolarStorm) Apply(p []interfaces.PlanetaryBody) {
	for _, body := range p {
		prop := body.Properties()
		// Adds some radiation and tempreture
		prop.SetTemperature(prop.GetTemperature() + 100*s.Intensity)
		prop.SetRadiation(prop.GetRadiation() + 100*s.Intensity)
	}
}
func (s *SolarStorm) Cease(p []interfaces.PlanetaryBody) {}
func (s *SolarStorm) Duration() time.Duration {
	return s.TimeLength
}
func (s *SolarStorm) NextEvent() interfaces.SystemEvent {
	// A storm can increase in intensity
	if rand.Intn(2) == 1 {
		return &SolarStorm{
			TimeLength: s.Duration(),
			Intensity:  s.Intensity * 2,
			Desc:       "The storm is increasing in intensity!",
		}
	}
	return &SolarStormCooling{
		TimeLength:    s.Duration() / 2,
		CoolingAmount: s.Intensity * 2,
	}
}
func (s *SolarStorm) Description() string {
	if s.Desc != "" {
		return s.Desc
	}
	return "there is a horrible storm!"
}
