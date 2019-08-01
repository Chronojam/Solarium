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

type TheWolfGamemode struct {
	// Use a map for faster lookups.
	// the key is a player's SecretID
	Players map[string]*proto.TheWolfGamePlayer

	// Who voted to lynch who?
	// key = playerId
	// value = number of votes
	LynchMap map[string]int
	// ActionLock to prevent race conditions if both players declare actions at the
	// same time.
	ActionLock sync.Mutex

	GameStartVotes  map[string]int
	GameStartedChan chan bool
	GameStarted     bool

	RoundStartChan chan bool

	// If its nighttime, only werewolves can vote to lynch
	IsNight bool

	EventStream chan *solarium.GameEvent
	Running     bool
}

func New() *TheWolfGamemode {
	return &TheWolfGamemode{
		LynchMap:        map[string]int{},
		GameStartedChan: make(chan bool, 2),
		GameStartVotes:  map[string]int{},
		GameStarted:     false,
		IsNight:         false,
		RoundStartChan:  make(chan bool, 2),
		EventStream:     make(chan *solarium.GameEvent),
		Players:         map[string]*proto.TheWolfGamePlayer{},
		Running:         true,
		ActionLock:      sync.Mutex{},
	}
}

func (t *TheWolfGamemode) Description() string {
	return ""
}
func (t *TheWolfGamemode) IsRunning() bool {
	return t.Running
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
	t.EventStream <- &solarium.GameEvent{
		Name: "A Player has joined!",
		Desc: "",
		TheWolfGame: &proto.TheWolfGameEvent{
			NewPlayer: &proto.TheWolfGameEvent_PlayerJoined{
				PlayerName: name,
			},
		},
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

	t.ActionLock.Lock()
	defer t.ActionLock.Unlock()
	// Start vote, only allow each player to vote once.
	// and only before the game has started. Otherwise ignore their request.
	if action.StartVote != nil && !t.GameStarted {
		t.GameStartVotes[pid] = 1
		if len(t.GameStartVotes) >= len(t.Players)/2 && len(t.Players) > RequiredPlayersForGame {
			// Assign werewolves, we need to alter the slice as we're iterating
			// (in case the same person is picked twice.)
			numWolves := len(t.Players) / 5
			for w := 0; w < numWolves; w++ {
				pindex := rand.Intn(len(t.Players))
				i := 0
				for _, p := range t.Players {
					if i == pindex {
						// if this person is already a werewolf, run
						// again
						if p.Role == proto.TheWolfGamePlayer_WEREWOLF {
							numWolves++
						}
						p.Role = proto.TheWolfGamePlayer_WEREWOLF
					}
					i++
				}
			}
			// GameStart condition
			t.GameStarted = true
			t.GameStartedChan <- true
			t.EventStream <- &solarium.GameEvent{
				Name: "The Game has started",
				Desc: "",
				TheWolfGame: &proto.TheWolfGameEvent{
					GameStart: &proto.TheWolfGameEvent_GameStarted{},
				},
			}
		}
		return nil
	}

	if action.StartVote != nil && t.GameStarted {
		// Game's already started, so ignore StartVote's
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
		if t.IsNight {
			if p.Role == proto.TheWolfGamePlayer_WEREWOLF && p.IsAlive {
				numRequired++
			}
		} else {
			if p.IsAlive {
				numRequired++
			}
		}
	}
	numSubmitted := 0
	for _, p := range t.LynchMap {
		numSubmitted += p
	}
	// Check if everyone has voted
	if numSubmitted == numRequired {
		t.RoundStartChan <- true
	}
	return nil
}

func (t *TheWolfGamemode) Simulate() {
	for {
		// Wait here until the game is ready to start
		// this is to avoid infinite for {} (and thus max cpu usage.)
		<-t.GameStartedChan
		log.Printf("GSPass")

		// Wait for each player to take an action before continuing.
		<-t.RoundStartChan
		log.Printf("RoundStarPAss")

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
			Name: "A Horrible Murder",
			Desc: message,
			TheWolfGame: &proto.TheWolfGameEvent{
				PlayerDied: &proto.TheWolfGameEvent_PlayerDeath{
					PlayerID:   p.ID,
					PlayerName: p.Name,
				},
			},
		}
		t.IsNight = !t.IsNight
		t.EventStream <- &solarium.GameEvent{
			Name: "TimeTransistion",
			Desc: "",
			TheWolfGame: &proto.TheWolfGameEvent{
				Transisition: &proto.TheWolfGameEvent_TimeTransistion{
					IsNight: t.IsNight,
				},
			},
		}

		// Check for win conditions
		if len(wPlayers) == len(vPlayers) {
			// Wolves Win
			t.EventStream <- &solarium.GameEvent{
				Name:       "Werewolf Victory",
				Desc:       fmt.Sprintf("The werewolves have overcome the town."),
				IsGameOver: true,
				TheWolfGame: &proto.TheWolfGameEvent{
					WolfVictory: &proto.TheWolfGameEvent_WerewolfVictory{},
				},
			}
			return
		}

		if len(wPlayers) == 0 {
			// Villagers win
			t.EventStream <- &solarium.GameEvent{
				Name:       "Villager Victory",
				Desc:       fmt.Sprintf("The werewolves have been purged from the town."),
				IsGameOver: true,
				TheWolfGame: &proto.TheWolfGameEvent{
					VillageVictory: &proto.TheWolfGameEvent_VillagerVictory{},
				},
			}
			return
		}

		t.LynchMap = map[string]int{}
		t.GameStartedChan <- true
	}
}
