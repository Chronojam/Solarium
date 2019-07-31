package thewolfgame

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	proto "github.com/chronojam/solarium/pkg/gamemodes/thewolfgame/proto"
	solarium "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// Use a map for faster lookups.
	// the key is a player's SecretID
	Players map[string]*proto.TheWolfGamePlayer

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
		LynchMap:        map[string]int{},
		GameStartedChan: make(chan bool),
		GameStartVotes:  map[string]int{},
		GameStarted:     false,
		IsNight:         false,
		RoundStartChan:  make(chan bool),
		EventStream:     make(chan *solarium.GameEvent),
		Players:         map[string]*proto.TheWolfGamePlayer{},
	}
}

func (t *TheWolfGamemode) Description() string {
	return ""
}

func (t *TheWolfGamemode) Status(pid, psecret string) (*solarium.GameStatusResponse, error) {
	// Strip role from response.
	players := []*proto.TheWolfGamePlayer{}
	if pid != "" && psecret != "" {
		p, ok := t.GetPlayer(pid, psecret, true)
		if !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Bad PlayerID/Secret")
		}
		players = []*proto.TheWolfGamePlayer{
			p,
		}
	} else {
		for _, p := range t.Players {
			players = append(players, &proto.TheWolfGamePlayer{
				ID:      p.ID,
				Name:    p.Name,
				IsAlive: p.IsAlive,
			})
		}
	}
	return &solarium.GameStatusResponse{
		TheWolfGame: &proto.TheWolfGameStatus{
			Players:   players,
			IsNight:   t.IsNight,
			IsStarted: t.GameStarted,
		},
	}, nil
}
func (t *TheWolfGamemode) NextEvent() *solarium.GameEvent {
	e := <-t.EventStream
	return e
}
func (t *TheWolfGamemode) Setup() {}
func (t *TheWolfGamemode) Join(name string) (*solarium.Player, error) {
	if t.GameStarted {
		// Cant join a game in progress.
		return nil, status.Errorf(codes.Unavailable, "Cannot join a game in progress!")
	}

	// Generate PID + Secret.
	pid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	secret, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	p := &solarium.Player{
		Name:   name,
		Secret: secret.String(),
		ID:     pid.String(),
	}
	t.Players[secret.String()] = &proto.TheWolfGamePlayer{
		ID:      p.ID,
		Name:    name,
		Role:    proto.TheWolfGamePlayer_VILLAGER,
		IsAlive: true,
	}
	return p, nil
}

func (t *TheWolfGamemode) GetPlayer(pid, secret string, validateSecret bool) (*proto.TheWolfGamePlayer, bool) {
	for s, p := range t.Players {
		if p.ID == pid {
			if validateSecret && s == secret {
				// Secret is valid
				return p, true
			}
			if validateSecret && s != secret {
				// Secret is invalid.
				return p, false
			}
			return p, true
		}
	}

	// PlayerID does not exist.
	return nil, false
}

func (t *TheWolfGamemode) PlayerDoAction(a interface{}, pid, secret string) error {
	action, ok := a.(*proto.TheWolfGameAction)
	if !ok {
		// This is an invalid action
		return status.Errorf(codes.InvalidArgument, "Invalid action for TheWolfGame")
	}
	player, ok := t.GetPlayer(pid, secret, true)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "PID or PSecret invalid")
	}
	// Start vote, only allow each player to vote once.
	if action.StartVote != nil {
		t.GameStartVotes[pid] = 1
		if len(t.GameStartVotes) >= len(t.Players)/2 && len(t.Players) > RequiredPlayersForGame {
			// Assign werewolves
			numWolves := len(t.Players) / 5
			for w := 0; w < numWolves; w++ {
				pindex := rand.Intn(len(t.Players))
				i := 0
				for _, p := range t.Players {
					if i == pindex {
						p.Role = proto.TheWolfGamePlayer_WEREWOLF
						log.Printf("%v is a werewolf", p.Name)
					}
					i++
				}
			}
			// GameStart condition
			t.GameStarted = true
			t.GameStartedChan <- true
		}
		return nil
	}
	// Dont allow anyone to do anything until the game has started
	if !t.GameStarted {
		return nil
	}
	if !player.IsAlive {
		// Cant vote if you are dead.
		return nil
	}
	if action.Vote == nil {
		return status.Errorf(codes.InvalidArgument, "Action was submitted, but no .Vote was nil.")
	}

	// Who are we voting for?
	pVote := action.Vote.PlayerId

	ActionLock.Lock()
	defer ActionLock.Unlock()
	if t.IsNight {
		// Only werewolves get to do something at night
		if player.Role == proto.TheWolfGamePlayer_WEREWOLF {
			t.LynchMap[pVote] = t.LynchMap[pVote] + 1
		}
	} else {
		t.LynchMap[pVote] = t.LynchMap[pVote] + 1
	}
	numRequired := 0
	for _, p := range t.Players {
		if p.IsAlive {
			numRequired++
		}
	}
	// Check if everyone has voted
	if len(t.LynchMap) == numRequired {
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
			if p.Role == proto.TheWolfGamePlayer_VILLAGER {
				vPlayers = append(vPlayers, &solarium.Player{
					Name: p.Name,
				})
			} else {
				wPlayers = append(wPlayers, &solarium.Player{
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

		p, ok := t.GetPlayer(toKillID, "", false)
		if !ok {
			log.Printf("Players elected to lynch a player who doesnt exist")
			log.Printf("Skipping round.")
			continue
		}

		// RIP
		message := fmt.Sprintf("The sun sets, the mob flys into a panic and heads to %v's house!; %v has been lynched by the mob!", p.Name, p.Name)
		if t.IsNight {
			message = fmt.Sprintf("A glorious new day rises, but %v hasnt turned up for church!", p.Name)
		}
		p.IsAlive = false
		t.EventStream <- &solarium.GameEvent{
			Name: "12 Hours Pass.",
			Desc: message,
			AffectedPlayers: []*solarium.Player{
				&solarium.Player{
					Name: p.Name,
				},
			},
			TheWolfGame: &proto.TheWolfGameEvent{
				PlayerDied: &proto.TheWolfGameEvent_PlayerDeath{},
			},
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
				TheWolfGame: &proto.TheWolfGameEvent{
					WolfVictory: &proto.TheWolfGameEvent_WerewolfVictory{},
				},
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
				TheWolfGame: &proto.TheWolfGameEvent{
					VillageVictory: &proto.TheWolfGameEvent_VillagerVictory{},
				},
			}
			return
		}

		t.IsNight = !t.IsNight
		t.LynchMap = map[string]int{}
		t.GameStartedChan <- true
	}
}
