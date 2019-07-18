package system

import (
	"log"
	"time"

	"github.com/chronojam/solarium/pkg/events"
	"github.com/chronojam/solarium/pkg/system/interfaces"
	"github.com/google/uuid"
)

type system struct {
	id     string
	bodies []interfaces.PlanetaryBody

	eventChan chan interfaces.SystemEvent
}

func init() {
	events.DEBUG_MODE = true
}

// New returns a new solar system
func New() *system {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	eventChan := make(chan interfaces.SystemEvent)
	go func() {
		eventChan <- events.RandomCalm()
	}()
	return &system{
		id:        uuid.String(),
		eventChan: eventChan,
	}
}

func (s *system) Simulate() {
	for {
		// Print out the state of the system
		s.NextEvent(s.eventChan)
	}
}

func (s *system) NextEvent(stream chan interfaces.SystemEvent) {
	e := <-stream
	log.Printf(e.Desc())
	e.Apply(s.bodies)
	go func() {
		t := time.NewTimer(e.Duration())
		<-t.C

		e.Cease(s.bodies)
		stream <- e.NextEvent()
	}()
}

// AddBody adds a new planetaryBody to this system
func (s *system) AddBody(p interfaces.PlanetaryBody) {
	s.bodies = append(s.bodies, p)
}