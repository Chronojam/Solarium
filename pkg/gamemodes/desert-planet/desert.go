package desert

import (
	"fmt"
	"log"
	"sync"

	proto "github.com/chronojam/solarium/pkg/gamemodes/desert/proto"
	solarium "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"
)

var (
	Descriptions = map[solarium.NewGameRequest_DifficultyLevel]string{
		// Easy
		solarium.NewGameRequest_EASY: `Stranded on a desert planet, there are a small amount of supplies scattered within the immediate vicinity of your landing area
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
		// Normal
		solarium.NewGameRequest_NORMAL: `Stranded on a desert planet, there are only a few supplies scattered within the immediate vicinity of your landing area
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
		// Hard
		solarium.NewGameRequest_HARD: `Stranded on a desert planet, none of your gear managed to survive the crash; 
		You must find more water & food to survive; Getting off the planet wont be easy - you'll need to gather components from the planet.
		Fuel will allow you to travel further in a day.`,
	}
)

// DesertGamemode; Players must gather appropriate resources
// in order to escape the planet, they must try and balance
// finding components with having enough water/food/shelter
type DesertGamemode struct {
	Difficulty solarium.NewGameRequest_DifficultyLevel
	GameStatus *proto.DesertPlanetStatus

	// The group's Score'
	Score int

	// The players
	Players []*solarium.Player

	// Event stream
	EventStream  chan *solarium.GameEvent
	RoundActions map[string]interface{}
}

func (d *DesertGamemode) Description() string {
	return Descriptions[d.Difficulty]
}
func (d *DesertGamemode) Status() *solarium.GameStatusResponse {
	return &solarium.GameStatusResponse{
		DesertPlanet: d.GameStatus,
	}
}

func New(difficulty solarium.NewGameRequest_DifficultyLevel) *DesertGamemode {
	return &DesertGamemode{
		Difficulty:   difficulty,
		EventStream:  make(chan *solarium.GameEvent),
		Score:        100 + int(difficulty)*100,
		RoundActions: map[string]interface{}{},
		GameStatus: &proto.DesertPlanetStatus{
			PlayerStatus: []*proto.DesertPlanetPlayerStatus{},
		},
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

	d.GameStatus.PlayerStatus = append(d.GameStatus.PlayerStatus, &proto.DesertPlanetPlayerStatus{
		PlayerID:     p.ID,
		PlayerName:   p.Name,
		Hunger:       3,
		Thirst:       3,
		Incapaciated: false,
		Status:       []string{},
	})

	d.Players = append(d.Players, p)
	return p, nil
}

func (d *DesertGamemode) FindStatusByPid(pid string) (*proto.DesertPlanetPlayerStatus, bool) {
	for _, s := range d.GameStatus.PlayerStatus {
		if s.PlayerID == pid {
			return s, true
		}
	}

	return nil, false
}

func (d *DesertGamemode) FindPlayerByPid(pid string) (*solarium.Player, bool) {
	for _, p := range d.Players {
		if p.ID == pid {
			return p, true
		}
	}

	return nil, false
}

var ActionLock = sync.Mutex{}

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
	ActionLock.Lock()
	defer ActionLock.Unlock()
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
		if len(d.Players) == 0 {
			log.Printf("Waiting for players to join")
			continue
		}
		// Wait for each player to take an action before continuing
		if len(d.RoundActions) != len(d.Players) {
			log.Printf("Waiting for all players to declare an action..")
			continue
		}
		// Resolve the round.
		d.ResolveRound()

		// Check for win & lose conditions
		incappedPlayers := 0
		for _, p := range d.GameStatus.PlayerStatus {
			if p.Incapaciated {
				incappedPlayers += 1
			}
		}
		// Everyone is down.
		if incappedPlayers == len(d.Players) {
			d.EventStream <- &solarium.GameEvent{
				Name:            "Failed",
				Desc:            fmt.Sprintf("All the players have died! Gameover!"),
				AffectedPlayers: d.Players,
				DesertPlanet: &proto.DesertPlanetEvent{
					DesertPlanetFailed: &proto.DesertPlanetFailed{
						Score: int32(d.Score),
					},
				},
			}
			break
		}

		if d.GameStatus.TargetComponents == d.GameStatus.Components {
			d.EventStream <- &solarium.GameEvent{
				Name:            "Succeeded",
				Desc:            fmt.Sprintf("The players managed to escape!"),
				AffectedPlayers: d.Players,
				DesertPlanet: &proto.DesertPlanetEvent{
					DesertPlanetSucceeded: &proto.DesertPlanetSucceeded{
						Score: int32(d.Score),
					},
				},
			}
			break
		}

		log.Printf("Resolved Round!")
	}
}

func (d *DesertGamemode) ResolveRound() {
	// See what state each player is in
	for _, p := range d.Players {
		status, _ := d.FindStatusByPid(p.ID)
		status.Hunger -= 1
		status.Thirst -= 1

		// You can only eat/drink if you are not incapacitated.
		if !status.Incapaciated {
			if status.Thirst <= 2 {
				// Try to take a drink.
				if d.GameStatus.Water >= 0 {
					d.GameStatus.Water -= 1
					status.Thirst += 2
				}
			}
			if status.Hunger <= 2 {
				// Try to get something to eat.
				if d.GameStatus.Food >= 0 {
					d.GameStatus.Food -= 1
					status.Hunger += 2
				}
			}
		} else {
			// If you are incapped, you cant take an action - so clear your action from the queue.
			delete(d.RoundActions, p.ID)
		}

		status.Status = []string{}
		switch {
		case status.Thirst <= 0:
			status.Incapaciated = true
			status.Status = append(status.Status, fmt.Sprintf("%v has collapsed due to thirst!", p.Name))
			d.Score -= 5
		case status.Thirst > 0 && status.Thirst < 3:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is extremely thirsty", p.Name))
			d.Score -= 1
		case status.Thirst >= 3:
			status.Incapaciated = false
			status.Status = append(status.Status, fmt.Sprintf("%v is perfectly hydrated", p.Name))
		}
		switch {
		case status.Hunger <= 0:
			status.Incapaciated = true
			status.Status = append(status.Status, fmt.Sprintf("%v has collapsed due to hunger!", p.Name))
			d.Score -= 5
		case status.Hunger > 0 && status.Hunger < 3:
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
	case solarium.NewGameRequest_EASY:
		d.GameStatus.Fuel = 2
		d.GameStatus.Water = 2
		d.GameStatus.Food = 2
		d.GameStatus.Components = 0
		d.GameStatus.TargetComponents = 3
	case solarium.NewGameRequest_NORMAL:
		d.GameStatus.Fuel = 1
		d.GameStatus.Water = 1
		d.GameStatus.Food = 1
		d.GameStatus.Components = 0
		d.GameStatus.TargetComponents = 5
	case solarium.NewGameRequest_HARD:
		d.GameStatus.Fuel = 0
		d.GameStatus.Water = 0
		d.GameStatus.Food = 0
		d.GameStatus.Components = 0
		d.GameStatus.TargetComponents = 10
	default:
		d.GameStatus.Fuel = 2
		d.GameStatus.Water = 2
		d.GameStatus.Food = 2
		d.GameStatus.Components = 0
		d.GameStatus.TargetComponents = 3
	}
}

func (d *DesertGamemode) NextEvent() *solarium.GameEvent {
	e := <-d.EventStream
	return e
}
