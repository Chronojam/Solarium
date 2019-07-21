package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chronojam/solarium/clients/discord/pkg/discord"
	"github.com/chronojam/solarium/clients/discord/pkg/handlers/newgame"
)

// https://discordapp.com/api/oauth2/authorize?client_id=602545365891285015&scope=bot&permissions=3088

func main() {
	d, err := discord.New(
		"!solarium",
		"NjAyNTQ1MzY1ODkxMjg1MDE1.XTSaPw.9o_ukUnYujMe8FL0TqGgZzTtkdk",
		"localhost:8443",
		"445700976582590497",
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	newgame.New(d)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	d.Run(sc)
}
