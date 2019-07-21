package discord

import (
	"context"
	"io"
	"log"
	"os"
	"strings"

	proto "github.com/chronojam/solarium/proto"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/grpc"
)

type Handler interface {
	Handle(sess *discordgo.Session, m *discordgo.MessageCreate)
}

type Bot struct {
	dg              *discordgo.Session
	responsePrefix  string
	Handlers        []Handler
	SolariumClient  proto.SolariumClient
	GlobalChannelID string
}

func New(responsePrefix, token, solariumAddr, globalChanName string) (*Bot, error) {
	// Handle discord connections
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	// Handle Solarium connections
	conn, err := grpc.Dial(solariumAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := proto.NewSolariumClient(conn)
	return &Bot{
		SolariumClient:  client,
		dg:              dg,
		responsePrefix:  responsePrefix,
		GlobalChannelID: globalChanName,
	}, nil
}

func (b *Bot) BroadcastGameEvents(channelId, gameId string) {
	stream, err := b.SolariumClient.GlobalUpdate(context.Background(), &proto.GameUpdateRequest{
		GameID: gameId,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	for {
		not, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		_, err = b.dg.ChannelMessageSend(channelId, not.Notification)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func (b *Bot) BroadcastToGlobalChan() {
	stream, err := b.SolariumClient.GlobalUpdate(context.Background(), &proto.GlobalUpdateRequest{})
	if err != nil {
		log.Fatalf("%v", err)
	}

	for {
		not, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		_, err = b.dg.ChannelMessageSend(b.GlobalChannelID, not.Notification)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func (b *Bot) AddHandler(h Handler) {
	b.dg.AddHandler(h.Handle)
}

func (b *Bot) ParseMessage(m *discordgo.MessageCreate) (bool, []string) {
	args := strings.Split(m.Content, " ")
	// Ignore anything that doesnt start with the global prefix.
	if args[0] != b.responsePrefix {
		return false, args
	}

	if len(args) == 1 {
		// Also put in help text here.
		return false, args
	}

	return true, args
}

func (b *Bot) Run(c chan os.Signal) {
	err := b.dg.Open()
	if err != nil {
		log.Fatalf("%v", err)
	}
	go b.BroadcastToGlobalChan()

	log.Printf("Bot is now running. Press CTRL-C to exit.")
	<-c

	b.dg.Close()
}
