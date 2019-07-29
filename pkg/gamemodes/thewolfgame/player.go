package thewolfgame

type PlayerRole int32

const (
	PlayerRole_Villager PlayerRole = 0
	PlayerRole_Werewolf PlayerRole = 1
)

type TheWolfGamePlayer struct {
	ID         string
	Name       string
	PlayerRole PlayerRole
	IsAlive    bool
}
