package events

import (
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

type SolarStormCooling struct {
	TimeLength    time.Duration
	CoolingAmount int
	Desc          string
}

func (s *SolarStormCooling) Apply(p []interfaces.PlanetaryBody) {}
func (s *SolarStormCooling) Cease(p []interfaces.PlanetaryBody) {}
func (s *SolarStormCooling) Duration() time.Duration {
	return s.TimeLength
}
func (s *SolarStormCooling) NextEvent() interfaces.SystemEvent {
	r := RandomCalm()
	r.Desc = "The storm is over"
	return r
}
func (s *SolarStormCooling) Description() string {
	if s.Desc != "" {
		return s.Desc
	}
	return "The storm is clearing"
}
