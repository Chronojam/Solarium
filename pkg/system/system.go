package system

import (
	"time"

	"github.com/chronojam/solarium/pkg/events"
	"github.com/chronojam/solarium/pkg/system/interfaces"
	"github.com/google/uuid"
)

type System struct {
	id     string
	bodies []interfaces.PlanetaryBody

	eventChan        chan interfaces.SystemEvent
	NotificationChan chan string
}

func init() {
	events.DEBUG_MODE = true
}

// New returns a new solar system
func New() *System {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	eventChan := make(chan interfaces.SystemEvent)
	go func() {
		eventChan <- events.RandomCalm()
	}()
	return &System{
		id:               uuid.String(),
		eventChan:        eventChan,
		NotificationChan: make(chan string),
	}
}

func (s *System) Simulate() {
	for {
		// Print out the state of the system
		s.NextEvent(s.eventChan)
	}
}

func (s *System) NextEvent(stream chan interfaces.SystemEvent) {
	e := <-stream
	s.NotificationChan <- e.Desc()
	e.Apply(s.bodies)
	go func() {
		t := time.NewTimer(e.Duration())
		<-t.C

		e.Cease(s.bodies)
		stream <- e.NextEvent()
	}()
}

// AddBody adds a new planetaryBody to this system
func (s *System) AddBody(p interfaces.PlanetaryBody) {
	s.bodies = append(s.bodies, p)
}
