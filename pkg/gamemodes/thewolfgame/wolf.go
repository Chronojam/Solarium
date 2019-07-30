package thewolfgame

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	proto "github.com/chronojam/solarium/pkg/gamemodes/thewolfgame/proto"
	solarium "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"
)

const (
	// Require at *least* 3 players
	RequiredPlayersForGame = 5
	WerewolvesPerPlayer    = 0.4
)

var (
	// ActionLock to prevent race conditions if both players declare actions at the
	// same time.
	ActionLock = sync.Mutex{}
)

type TheWolfGamemode struct {
	GameStatus *proto.TheWolfGameStatus
	// Use a map for faster lookups.
	Players map[string]*TheWolfGamePlayer

	// Who voted to lynch who?
	// key = playerId
	// value = number of votes
	LynchMap map[string]int

	GameStartVotes  map[string]int
	GameStartedChan chan bool
	GameStarted     bool

	RoundStartChan chan bool

	// If its nighttime, only werewolves can vote to lynch
	IsNight bool

	EventStream chan *solarium.GameEvent
}

func New() *TheWolfGamemode {
	return &TheWolfGamemode{
		GameStatus:      &proto.TheWolfGameStatus{},
		LynchMap:        map[string]int{},
		GameStartedChan: make(chan bool),
		GameStartVotes:  map[string]int{},
		GameStarted:     false,
		IsNight:         false,
		RoundStartChan:  make(chan bool),
		EventStream:     make(chan *solarium.GameEvent),
		Players:         map[string]*TheWolfGamePlayer{},
	}
}

func (t *TheWolfGamemode) Description() string {
	return ""
}

func (t *TheWolfGamemode) Status() *solarium.GameStatusResponse {
	return &solarium.GameStatusResponse{
		TheWolfGame: t.GameStatus,
	}
}
func (t *TheWolfGamemode) NextEvent() *solarium.GameEvent {
	e := <-t.EventStream
	return e
}
func (t *TheWolfGamemode) Setup() {
	t.GameStatus.IsNight = t.IsNight
}
func (t *TheWolfGamemode) Join(name string) (*solarium.Player, error) {
	if t.GameStarted {
		// Cant join a game in progress.
		return nil, nil
	}
	pid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	p := &solarium.Player{
		Name: name,
		ID:   pid.String(),
	}
	// Setup the player.
	t.Players[p.ID] = &TheWolfGamePlayer{
		ID:         p.ID,
		Name:       name,
		PlayerRole: PlayerRole_Villager,
		IsAlive:    true,
	}
	return p, nil
}

func (t *TheWolfGamemode) PlayerDoAction(req *solarium.DoActionRequest) error {
	pid := req.PlayerID
	player, ok := t.Players[pid]
	if !ok {
		// player doesnt exist,
		return nil
	}
	if req.TheWolfGame == nil {
		return nil
	}
	// Start vote, only allow each player to vote once.
	if req.TheWolfGame.StartVote != nil {
		t.GameStartVotes[pid] = 1
		if len(t.GameStartVotes) >= len(t.Players)/2 && len(t.Players) > RequiredPlayersForGame {
			// Assign werewolves
			numWolves := len(t.Players) / 5
			for w := 0; w < numWolves; w++ {
				pindex := rand.Intn(len(t.Players))
				i := 0
				for _, p := range t.Players {
					if i == pindex {
						p.PlayerRole = PlayerRole_Werewolf
						log.Printf("%v is a werewolf", p.Name)
					}
					i++
				}
			}
			// GameStart condition
			t.GameStarted = true
			t.GameStartedChan <- true

			// Send a temporary gameevent for now telling any listening client
			// who the werewolves are.
			// who are the wolves?
			players := []*proto.TheWolfGameStatusPlayer{}
			for _, p := range t.Players {
				players = append(players, &proto.TheWolfGameStatusPlayer{
					Name: p.Name,
					Role: int32(p.PlayerRole),
				})
			}
			t.EventStream <- &solarium.GameEvent{
				Name: "The Werewolves have been selected",
				Desc: "",
				TheWolfGame: &proto.TheWolfGameEvent{
					Players: players,
				},
			}

			t.GameStatus.Players = players
		}
		return nil
	}
	// Dont allow anyone to do anything until the game has started
	if !t.GameStarted {
		return nil
	}
	if req.TheWolfGame.Vote == nil {
		return nil
	}

	// Who are we voting for?
	pVote := req.TheWolfGame.Vote.PlayerId

	ActionLock.Lock()
	defer ActionLock.Unlock()
	if t.IsNight {
		// Only werewolves get to do something at night
		if player.PlayerRole == PlayerRole_Werewolf {
			t.LynchMap[pVote] = t.LynchMap[pVote] + 1
		}
	} else {
		t.LynchMap[pVote] = t.LynchMap[pVote] + 1
	}
	// Check if everyone has voted
	if len(t.LynchMap) == len(t.Players) {
		t.RoundStartChan <- true
	}
	return nil
}

func (t *TheWolfGamemode) Simulate() {
	for {
		// Wait here until the game is ready to start
		// this is to avoid infinite for {} (and thus max cpu usage.)
		<-t.GameStartedChan

		// Wait for each player to take an action before continuing.
		<-t.RoundStartChan

		// The Grouping of players who are still alive
		// expressed as solarium.Players.
		vPlayers := []*solarium.Player{}
		wPlayers := []*solarium.Player{}
		for _, p := range t.Players {
			if !p.IsAlive {
				continue
			}
			if p.PlayerRole == PlayerRole_Villager {
				vPlayers = append(vPlayers, &solarium.Player{
					Name: p.Name,
				})
			} else {
				wPlayers = append(vPlayers, &solarium.Player{
					Name: p.Name,
				})
			}
		}

		// Decide who gets killed
		toKillID := ""
		highest := 0
		for id, count := range t.LynchMap {
			// In the event of a draw, the last one
			// gets lynched, sorry!
			if count >= highest {
				highest = count
				toKillID = id
			}
		}

		// RIP
		allPlayers := append(vPlayers, wPlayers...)
		supposedEvilDoer := ""
		for supposedEvilDoer != "" {
			// Pick a randomer who isnt the one who got killed
			i := rand.Intn(len(allPlayers))
			supposedEvilDoer = t.Players[allPlayers[i].ID].Name
		}
		message := fmt.Sprintf("The sun sets, the mob flys into a panic and heads to %v's house!; %v has been lynched by the mob!", t.Players[toKillID].Name, t.Players[toKillID].Name)
		if t.IsNight {
			message = fmt.Sprintf("A glorious new day rises, but %v hasnt turned up for church! %v goes to investigate and finds them dead in bed!", t.Players[toKillID].Name, supposedEvilDoer)
		}
		t.Players[toKillID].IsAlive = false
		t.EventStream <- &solarium.GameEvent{
			Name: "12 Hours Pass.",
			Desc: message,
			AffectedPlayers: []*solarium.Player{
				&solarium.Player{
					Name: t.Players[toKillID].Name,
				},
			},
			TheWolfGame: &proto.TheWolfGameEvent{},
		}

		// Check for win conditions
		if len(wPlayers) == len(vPlayers) {
			// Wolves Win
			t.EventStream <- &solarium.GameEvent{
				Name:             "Werewolf Victory",
				Desc:             fmt.Sprintf("The werewolves have overcome the town."),
				InitatingPlayers: wPlayers,
				AffectedPlayers:  vPlayers,
				IsGameOver:       true,
				TheWolfGame:      &proto.TheWolfGameEvent{},
			}
			return
		}

		if len(wPlayers) == 0 {
			// Villagers win
			t.EventStream <- &solarium.GameEvent{
				Name:             "Villager Victory",
				Desc:             fmt.Sprintf("The werewolves have been purged from the town."),
				InitatingPlayers: vPlayers,
				AffectedPlayers:  wPlayers,
				IsGameOver:       true,
				TheWolfGame:      &proto.TheWolfGameEvent{},
			}
			return
		}

		t.IsNight = !t.IsNight
		t.GameStatus.IsNight = t.IsNight
		t.LynchMap = map[string]int{}
		t.GameStartedChan <- true
		log.Printf("There are %v Villagers, and %v Wolves", len(vPlayers), len(wPlayers))
	}
}
