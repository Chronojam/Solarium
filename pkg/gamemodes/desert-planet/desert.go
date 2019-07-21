package desert

import (
	"fmt"
	"log"

	"github.com/chronojam/solarium/pkg/game/interfaces"
	components "github.com/chronojam/solarium/pkg/gamemodes/desert-planet/jobs/gather-components"
	food "github.com/chronojam/solarium/pkg/gamemodes/desert-planet/jobs/gather-food"
	water "github.com/chronojam/solarium/pkg/gamemodes/desert-planet/jobs/gather-water"
	idle "github.com/chronojam/solarium/pkg/gamemodes/desert-planet/jobs/idle"
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player"
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player/conditions"
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player/conditions/hunger"
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet/player/conditions/thirst"
	"github.com/google/uuid"
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

	PlayerActions = map[string]func(p *player.Player, d *DesertGamemode){
		"me:GetWater": func(p *player.Player, d *DesertGamemode) {
			p.Job = water.New(false)
			d.SendEvent(fmt.Sprintf("%v is now gathering water!", p.Name))
		},
		"me:GetFood": func(p *player.Player, d *DesertGamemode) {
			p.Job = food.New(false)
			d.SendEvent(fmt.Sprintf("%v is now gathering food!", p.Name))
		},
		"me:GetComponents": func(p *player.Player, d *DesertGamemode) {
			p.Job = components.New(false)
			d.SendEvent(fmt.Sprintf("%v is now gathering components!", p.Name))
		},
		"me:GetWaterWithFuel": func(p *player.Player, d *DesertGamemode) {
			if d.Fuel > 0 {
				p.Job = water.New(true)
				d.Fuel--
				d.SendEvent(fmt.Sprintf("%v has taken a truck to gather some water!", p.Name))
				return
			}
			p.Job = water.New(false)
			d.SendEvent(fmt.Sprintf("%v is now gathering water!", p.Name))
		},
		"me:GetFoodWithFuel": func(p *player.Player, d *DesertGamemode) {
			if d.Fuel > 0 {
				p.Job = food.New(true)
				d.Fuel--
				d.SendEvent(fmt.Sprintf("%v has taken a truck to gather some food!", p.Name))
				return
			}
			p.Job = food.New(false)
			d.SendEvent(fmt.Sprintf("%v is now gathering food!", p.Name))
		},
		"me:GetComponentsWithFuel": func(p *player.Player, d *DesertGamemode) {
			if d.Fuel > 0 {
				p.Job = components.New(true)
				d.Fuel--
				d.SendEvent(fmt.Sprintf("%v has taken a truck to gather some components!", p.Name))
				return
			}
			p.Job = components.New(true)
			d.SendEvent(fmt.Sprintf("%v is now gathering components!", p.Name))
		},
	}
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

func New(difficulty int) *DesertGamemode {
	return &DesertGamemode{
		Difficulty: difficulty,
	}
}

func (d *DesertGamemode) Join(name string) (interfaces.Player, error) {
	pid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	p := &player.Player{
		Name: name,
		ID:   pid.String(),
		Job:  &idle.Idle{},
	}

	d.Players = append(d.Players, p)
	p.Conditions = append(p.Conditions, thirst.New())
	p.Conditions = append(p.Conditions, hunger.New())

	return p, nil
}

func (d *DesertGamemode) PlayerDoAction(pid, action string) error {
	for _, player := range d.Players {
		if player.Id() == pid {
			if val, ok := PlayerActions[action]; ok {
				val(player, d)
			}
		}
	}

	return nil
}

func (d *DesertGamemode) Setup() {
	switch d.Difficulty {
	case 1:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
		d.SendEvent(Easy.Description)
	case 2:
		d.Fuel = 1
		d.Water = 1
		d.Food = 1
		d.Components = 0
		d.TargetComponents = 5
		d.SendEvent(Normal.Description)
	case 3:
		d.Fuel = 0
		d.Water = 0
		d.Food = 0
		d.Components = 0
		d.TargetComponents = 10
		d.SendEvent(Hard.Description)
	default:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
		d.SendEvent(Easy.Description)
	}
}

func (d *DesertGamemode) NextEvent() interfaces.GameEvent {
	e := <-d.EventStream
	log.Printf("Fetched %v from estream", e.Description)
	return e
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
				d.SendEvent(fmt.Sprintf("%v units of water remaining", d.Water))
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
				d.SendEvent(fmt.Sprintf("%v has fetched %v units of water; there is now %v units of water", p.Name, v.Amount(), d.Water))
				p.Job = water.New(false) //idle.New(false)
			}
		case *food.Food:
			if v.Done() {
				d.Food += v.Amount()
				d.SendEvent(fmt.Sprintf("%v has fetched %v units of food", p.Name, v.Amount()))
				p.Job = idle.New(false)
			}
		case *components.Components:
			if v.Done() {
				d.Components += v.Amount()
				d.SendEvent(fmt.Sprintf("%v has fetched %v units of components", p.Name, v.Amount()))
				p.Job = idle.New(false)
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

	//d.SendEvent(fmt.Sprintf("GameState: \nWater: %v\nFood: %v\nFuel: %v\nComponents: %v\n", d.Water, d.Food, d.Fuel, d.Components))
}

func (d *DesertGamemode) SendEvent(des string) {
	go func() {
		//log.Printf("Sending Event: %v", des)
		d.EventStream <- interfaces.GameEvent{
			Description: des,
		}
	}()
}
