package desert

import (
	"fmt"
	"log"

	proto "github.com/chronojam/solarium/pkg/gamemodes/desert/proto"
	solarium "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"
)

var (
	Descriptions = map[int]string{
		// Easy
		0: `Stranded on a desert planet, there are a small amount of supplies scattered within the immediate vicinity of your landing area
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
		// Normal
		1: `Stranded on a desert planet, there are only a few supplies scattered within the immediate vicinity of your landing area
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
		// Hard
		2: `Stranded on a desert planet, none of your gear managed to survive the crash; 
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
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
	Score            int

	// The players
	Players []*solarium.Player
	// Haters will say you should store these on the
	// player object itself, but fight me?
	PlayerStatus map[string]*PlayerStatus

	// Event stream
	EventStream  chan *solarium.GameEvent
	RoundActions map[string]interface{}
}

func (d *DesertGamemode) Description() string {
	return Descriptions[d.Difficulty]
}

func New(difficulty int) *DesertGamemode {
	return &DesertGamemode{
		Difficulty:  difficulty,
		EventStream: make(chan *solarium.GameEvent),
		Score:       100 + difficulty*100,
	}
}

func (d *DesertGamemode) Join(name string) (*solarium.Player, error) {
	pid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	p := &solarium.Player{
		Name: name,
		ID:   pid.String(),
	}

	d.Players = append(d.Players, p)
	return p, nil
}

func (d *DesertGamemode) FindPlayerByPid(pid string) (*solarium.Player, bool) {
	for _, p := range d.Players {
		if p.ID == pid {
			return p, true
		}
	}

	return nil, false
}

func (d *DesertGamemode) PlayerDoAction(req *solarium.DoActionRequest) error {
	pid := req.PlayerID
	if _, ok := d.FindPlayerByPid(pid); !ok {
		// Invalid PID
		return nil
	}

	if req.DesertPlanet == nil {
		// Tried to do something, but didnt pass an action
		// or the action wasnt part of the right namespace.
		return nil
	}
	if req.DesertPlanet.GatherWater != nil {
		d.RoundActions[pid] = req.DesertPlanet.GatherWater
		return nil
	}
	if req.DesertPlanet.GatherFood != nil {
		d.RoundActions[pid] = req.DesertPlanet.GatherFood
		return nil
	}
	if req.DesertPlanet.GatherComponents != nil {
		d.RoundActions[pid] = req.DesertPlanet.GatherComponents
		return nil
	}

	// Should never get here, but here we are?
	log.Printf("Reached impossible in desert-Planet PlayerDoAction()")
	return nil
}

func (d *DesertGamemode) Simulate() {
	// As long as we've got players in the game.
	for {
		// For now, dont do anything while we wait for players.
		if len(d.Players) < 0 {
			continue
		}
		// Wait for each player to take an action before continuing
		if len(d.RoundActions) != len(d.Players) {
			continue
		}
		// Resolve the round.
		d.ResolveRound()

		// Check for win & lose conditions
		incappedPlayers := 0
		for _, p := range d.PlayerStatus {
			if p.Incapaciated {
				incappedPlayers += 1
			}
		}
		// Everyone is down.
		if incappedPlayers == len(d.Players)+1 {
			d.EventStream <- &solarium.GameEvent{
				Name:            "Failed",
				Desc:            fmt.Sprintf("All the players have died! Gameover!"),
				AffectedPlayers: d.Players,
				DesertPlanet: &proto.DesertPlanetEvent{
					DesertPlanetFailed: &proto.DesertPlanetFailed{},
				},
			}
			break
		}

		if d.TargetComponents == d.Components {
			d.EventStream <- &solarium.GameEvent{
				Name:            "Succeeded",
				Desc:            fmt.Sprintf("The players managed to escape!"),
				AffectedPlayers: d.Players,
				DesertPlanet: &proto.DesertPlanetEvent{
					DesertPlanetSucceeded: &proto.DesertPlanetSucceeded{},
				},
			}
			break
		}
	}
}

func (d *DesertGamemode) ResolveRound() {
	// See what state each player is in
	for _, p := range d.Players {
		status, _ := d.PlayerStatus[p.ID]
		status.Hunger -= 1
		status.Thirst -= 1

		// You can only eat/drink if you are not incapacitated.
		if !status.Incapaciated {
			if status.Thirst <= 2 {
				// Try to take a drink.
				if d.Water >= 0 {
					d.Water -= 1
					status.Thirst += 2
				}
			}
			if status.Hunger <= 2 {
				// Try to get something to eat.
				if d.Food >= 0 {
					d.Food -= 1
					status.Hunger += 2
				}
			}
		}

		status.Status = []string{}
		switch {
		case status.Thirst < 0:
			status.Incapaciated = true
			status.Status = append(status.Status, fmt.Sprintf("%v has collapsed due to thirst!", p.Name))
			d.Score -= 5
		case status.Thirst == 1:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is extremely thirsty", p.Name))
			d.Score -= 1
		case status.Thirst >= 2:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is perfectly hydrated", p.Name))
		}
		switch {
		case status.Hunger < 0:
			status.Incapaciated = true
			status.Status = append(status.Status, fmt.Sprintf("%v has collapsed due to hunger!", p.Name))
			d.Score -= 5
		case status.Hunger == 1:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is extremely hungry", p.Name))
			d.Score -= 1
		case status.Hunger >= 2:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is well fed", p.Name))
		}
	}

	// Allow all our players to do their actions now
	for pid, action := range d.RoundActions {
		player, _ := d.FindPlayerByPid(pid)
		status, _ := d.PlayerStatus[pid]
		if status.Incapaciated {
			// You get to do nothing because you are incapped.
			break
		}
		switch action.(type) {
		case *proto.DesertPlanetGatherWater:
			d.gatherWater(player)
		case *proto.DesertPlanetGatherFood:
			d.gatherFood(player)
		case *proto.DesertPlanetGatherComponents:
			d.gatherComponent(player)
		default:
			// Whatever we just got we dont reckonize.
		}
	}
	// Empty the buffer
	d.RoundActions = map[string]interface{}{}
}

func (d *DesertGamemode) Setup() {
	switch d.Difficulty {
	case 1:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
	case 2:
		d.Fuel = 1
		d.Water = 1
		d.Food = 1
		d.Components = 0
		d.TargetComponents = 5
	case 3:
		d.Fuel = 0
		d.Water = 0
		d.Food = 0
		d.Components = 0
		d.TargetComponents = 10
	default:
		d.Fuel = 2
		d.Water = 2
		d.Food = 2
		d.Components = 0
		d.TargetComponents = 3
	}
}

func (d *DesertGamemode) NextEvent() *solarium.GameEvent {
	e := <-d.EventStream
	return e
}
