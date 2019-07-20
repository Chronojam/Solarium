package desert

import (
	"github.com/chronojam/Solarium/pkg/game/interfaces"
	components "github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/jobs/gather-components"
	food "github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/jobs/gather-food"
	water "github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/jobs/gather-water"
	idle "github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/jobs/idle"
	"github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/player"
	"github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/player/conditions"
	"github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/player/conditions/hunger"
	"github.com/chronojam/Solarium/pkg/gamemodes/desert-planet/player/conditions/thirst"
)

var (
	Easy = interfaces.GameEvent{
		Description: `
Stranded on a desert planet, there are a small amount of supplies scattered within the immediate vicinity of your landing area
You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
Fuel will allow you to travel further in a day.
`}
	Normal = interfaces.GameEvent{
		Description: `
Stranded on a desert planet, there are only a few supplies scattered within the immediate vicinity of your landing area
You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
Fuel will allow you to travel further in a day.
`}
	Hard = interfaces.GameEvent{
		Description: `
Stranded on a desert planet, none of your gear managed to survive the crash; 
You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
Fuel will allow you to travel further in a day.
`}
)

// DesertGamemode; Players must gather appropriate resources
// in order to escape the planet, they must try and balance
// finding components with having enough water/food/shelter
type DesertGamemode struct {
	Difficulty int

	// The group inventory!
	Water            int
	Food             int
	Fuel             int
	Components       int
	TargetComponents int

	// The players
	Players     []*player.Player
	Description string

	// Event stream
	EventStream chan interfaces.GameEvent
}

func New(difficulty int, players []*player.Player) *DesertGamemode {
	return &DesertGamemode{
		Difficulty: difficulty,
		Players:    players,
	}
}

func (d *DesertGamemode) Setup() {
	switch d.Difficulty {
	case 1:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
		d.EventStream <- Easy
	case 2:
		d.Fuel = 1
		d.Water = 1
		d.Food = 1
		d.Components = 0
		d.TargetComponents = 5
		d.EventStream <- Normal
	case 3:
		d.Fuel = 0
		d.Water = 0
		d.Food = 0
		d.Components = 0
		d.TargetComponents = 10
		d.EventStream <- Hard
	default:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
		d.EventStream <- Easy
	}

	for _, p := range d.Players {
		// Make everyone thirsty.
		p.Conditions = append(p.Conditions, thirst.New())
	}
}

func (d *DesertGamemode) NextEvent() interfaces.GameEvent {
	return <-d.EventStream
}

// Do a single simulation 'tick'
func (d *DesertGamemode) Simulate() {
	playersWhoDied := []int{}
	for i, p := range d.Players {
		skipMe := false
		// Update all our conditions
		for _, c := range p.Conditions {
			// eat food/drink water first.
			switch v := c.(type) {
			case *thirst.Thirst:
				if d.Water == 0 {
					d.SendEvent(p.Name + " tried to take a drink of water, but there was none!")
					break
				}
				d.Water--
				v.SetValue(100)
			case *hunger.Hunger:
				if d.Food == 0 {
					d.SendEvent(p.Name + " tried to grab a bite to eat, but there was none!")
					break
				}
				d.Food--
				v.SetValue(100)
			}

			c.Simulate()

			// Decide who died.
			if c.State() == conditions.Fatal {
				playersWhoDied = append(playersWhoDied, i)
				d.SendEvent("Player: " + p.Name + " " + c.FatalText())
				skipMe = true
				continue
			}
		}
		if skipMe {
			continue
		}

		// Everyone continues to do their jobs.
		p.Job.Simulate()
		switch v := p.Job.(type) {
		case *water.Water:
			if v.Done() {
				d.Water += v.Amount()
			}
		case *food.Food:
			if v.Done() {
				d.Food += v.Amount()
			}
		case *components.Components:
			if v.Done() {
				d.Components += v.Amount()
			}
		case *idle.Idle:
			d.SendEvent(p.Name + " decides to laze around the camp for today")
		}
	}

	// Kill off anyone who died, remove them from the game
	for _, p := range playersWhoDied {
		d.Players = append(d.Players[:p], d.Players[p+1:]...)
	}

	// Lose Condition
	if len(d.Players) == 0 {
		d.SendEvent("Everyone died")
		return
	}

	// Win Condition
	if d.Components == d.TargetComponents {
		// You win!
		d.SendEvent("You got all the components! Nice one!")
		return
	}
}

func (d *DesertGamemode) SendEvent(des string) {
	go func() {
		d.EventStream <- interfaces.GameEvent{
			Description: des,
		}
	}()
}
