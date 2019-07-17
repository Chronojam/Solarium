package system

import (
	"math/rand"
	"testing"
	"time"

	"github.com/chronojam/solarium/pkg/events"
	"github.com/chronojam/solarium/pkg/system/interfaces"

	"github.com/chronojam/solarium/pkg/properties"

	"github.com/chronojam/solarium/pkg/planet"
)

func TestSimulation(t *testing.T) {
	eventChan := make(chan interfaces.SystemEvent)
	go func() {
		eventChan <- &events.SolarStorm{
			TimeLength: time.Second,
			Intensity:  rand.Intn(4),
		}
	}()
	s := &system{
		id:        "foobar",
		eventChan: eventChan,
	}
	b := planet.New(&planet.Properties{
		properties.KeyTemperature: 100,
		properties.KeyRadiation:   120,
	})
	s.AddBody(b)
	s.NextEvent()

	if b.Properties().GetInteger(properties.KeyTemperature) == 100 {
		t.Fail()
	}
	if b.Properties().GetInteger(properties.KeyRadiation) == 120 {
		t.Fail()
	}
}
