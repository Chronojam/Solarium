package desert

import (
	"fmt"
	"math/rand"

	proto "github.com/chronojam/solarium/pkg/gamemodes/desert/proto"
	solarium "github.com/chronojam/solarium/proto"
)

func (d *DesertGamemode) gatherWater(player *solarium.Player) {
	// How much water did we gather?, 1-4
	q := rand.Intn(3) + 1
	d.GameStatus.Water += int32(q)
	d.EventStream <- &solarium.GameEvent{
		Name: "Gathered Water",
		Desc: fmt.Sprintf("%s gathered %v units of water", player.Name, q),
		InitatingPlayers: []*solarium.Player{
			player,
		},
		DesertPlanet: &proto.DesertPlanetEvent{
			DesertPlanetGatheredWater: &proto.DesertPlanetGatheredWater{
				Quantity: int32(q),
			},
		},
	}
}

func (d *DesertGamemode) gatherFood(player *solarium.Player) {
	// How much food did we gather?, 1-4
	q := rand.Intn(3) + 1
	d.GameStatus.Food += int32(q)
	d.EventStream <- &solarium.GameEvent{
		Name: "Gathered Food",
		Desc: fmt.Sprintf("%s gathered %v units of food", player.Name, q),
		InitatingPlayers: []*solarium.Player{
			player,
		},
		DesertPlanet: &proto.DesertPlanetEvent{
			DesertPlanetGatheredFood: &proto.DesertPlanetGatheredFood{
				Quantity: int32(q),
			},
		},
	}
}

func (d *DesertGamemode) gatherComponent(player *solarium.Player) {
	// How much compoonents did we gather?, 1-2
	q := rand.Intn(1) + 1
	d.GameStatus.Components += int32(q)
	d.EventStream <- &solarium.GameEvent{
		Name: "Gathered Components",
		Desc: fmt.Sprintf("%s gathered %v units of components", player.Name, q),
		InitatingPlayers: []*solarium.Player{
			player,
		},
		DesertPlanet: &proto.DesertPlanetEvent{
			DesertPlanetGatheredComponents: &proto.DesertPlanetGatheredComponents{
				Quantity: int32(q),
			},
		},
	}
}
