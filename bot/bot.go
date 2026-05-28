package bot

import (
	"fmt"
	"log"

	"TradeBot/handlers"

	"github.com/bwmarrin/discordgo"
)

func Start(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents = discordgo.IntentsGuilds

	dg.AddHandler(handlers.InteractionHandler)

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		fmt.Println("RAW interaction received")
	})

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Bot started")

	return dg, nil
}
