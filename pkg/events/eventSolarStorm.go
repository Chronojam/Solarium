package events

import (
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/chronojam/solarium/pkg/planet/effects"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

// SolarStorm brews up a horrid storm, increasing tempreture and radiation
// on the body
type SolarStorm struct {
	TimeLength  time.Duration
	Intensity   int
	Description string

	effectUUID string
}

func (s *SolarStorm) Apply(p []interfaces.PlanetaryBody) {
	u, _ := uuid.NewRandom()
	s.effectUUID = u.String()
	for _, body := range p {
		body.AddEffect(&effects.SolarStorm{
			Identifier: s.effectUUID,
			Intensity:  s.Intensity,
		})
	}
}
func (s *SolarStorm) Cease(p []interfaces.PlanetaryBody) {
	for _, body := range p {
		body.RemoveEffect(&effects.SolarStorm{
			Identifier: s.effectUUID,
			Intensity:  s.Intensity,
		})
	}
}
func (s *SolarStorm) Duration() time.Duration {
	return s.TimeLength
}
func (s *SolarStorm) NextEvent() interfaces.SystemEvent {
	// A storm can increase in intensity
	if rand.Intn(2) == 1 {
		return &SolarStorm{
			TimeLength:  s.Duration(),
			Intensity:   s.Intensity * 2,
			Description: "The storm is increasing in intensity!",
		}
	}
	return RandomCalm()
}

func (s *SolarStorm) Desc() string {
	if s.Description != "" {
		return s.Description
	}
	return "A horrible storm is clouding the system."
}
