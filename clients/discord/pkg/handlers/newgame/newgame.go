package newgame

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/chronojam/solarium/clients/discord/pkg/discord"
	proto "github.com/chronojam/solarium/proto"
)

var (
	flagSet    = flag.NewFlagSet("newgame", flag.ContinueOnError)
	difficulty = flagSet.Int("difficulty", 1, "The difficulty of the game")
	gamemode   = flagSet.String("gamemode", "Desert-Planet", "The gamemode to play")
)

type NewGame struct {
	parent *discord.Bot
}

func New(b *discord.Bot) {
	n := &NewGame{
		parent: b,
	}
	b.AddHandler(n)
}

func (n *NewGame) Handle(sess *discordgo.Session, m *discordgo.MessageCreate) {
	ok, args := n.parent.ParseMessage(m)
	if !ok {
		return
	}
	if m.Author.ID == sess.State.User.ID {
		return
	}
	if args[1] != "newgame" {
		return
	}
	if err := flagSet.Parse(args[1:]); err != nil {
		return
	}

	_, err = sess.GuildChannelCreate("445700976582590495", resp.GameID, discordgo.ChannelTypeGuildText)
	if err != nil {
		log.Printf("%v", err)
		sess.ChannelMessageSend(n.parent.GlobalChannelID, fmt.Sprintf("%v", err))
		return
	}
	sess.ChannelMessageSend(n.parent.GlobalChannelID, fmt.Sprintf("#%v", resp.GameID))

	resp, err := n.parent.SolariumClient.NewGame(context.Background(), &proto.NewGameRequest{
		Gamemode:   *gamemode,
		Difficulty: int64(*difficulty),
	})
	if err != nil {
		sess.ChannelMessageSend(n.parent.GlobalChannelID, fmt.Sprintf("%v", err))
		return
	}
}
