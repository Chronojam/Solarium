package game

import (
	"time"

	"github.com/chronojam/solarium/pkg/events"
	"github.com/chronojam/solarium/pkg/game/interfaces"
	"github.com/google/uuid"
)

type player interface {
	ID() string
}

// Game is the overarcing container for the gamemode and
// players.
type Game struct {
	id      string
	players []player

	eventChan        chan interfaces.GameEvent
	NotificationChan chan string
}

func init() {
	events.DEBUG_MODE = true
}

// New returns a new solar Game
func New() *Game {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	eventChan := make(chan interfaces.GameEvent)
	go func() {
		eventChan <- events.RandomCalm()
	}()
	return &Game{
		id:               uuid.String(),
		eventChan:        eventChan,
		NotificationChan: make(chan string),
	}
}

func (s *Game) Simulate() {
	for {
		// Print out the state of the Game
		s.NextEvent(s.eventChan)
	}
}

func (s *Game) NextEvent(stream chan interfaces.GameEvent) {
	// Get the next event from the stream.
	e := <-stream
	// Send the description of the event to the notification channel.
	s.NotificationChan <- e.Desc()
	e.Apply(s.bodies)
	go func() {
		t := time.NewTimer(e.Duration())
		<-t.C

		e.Cease(s.bodies)
		stream <- e.NextEvent()
	}()
}

// AddBody adds a new planetaryBody to this Game
func (s *Game) AddBody(p interfaces.PlanetaryBody) {
	s.bodies = append(s.bodies, p)
}
