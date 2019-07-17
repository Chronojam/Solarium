package events

import (
	"math/rand"
	"time"

	"github.com/chronojam/solarium/pkg/system/interfaces"
)

const ()

func RandomEvent() interfaces.SystemEvent {
	// Generate a duration
	t := rand.Intn(360)
	duration := time.Second * time.Duration(t)
	if DEBUG_MODE {
		duration = time.Second
	}

	switch rand.Intn(2) {
	// SolarStorm
	case 0:
		return &SolarStorm{
			TimeLength: duration,
			Intensity:  rand.Intn(4),
		}
	default:
		return RandomCalm()
	}
}

func RandomCalm() *Calm {
	duration := time.Duration(rand.Intn(30)) * time.Second
	if DEBUG_MODE {
		duration = time.Second
	}
	return &Calm{
		TimeLength: duration,
	}
}
