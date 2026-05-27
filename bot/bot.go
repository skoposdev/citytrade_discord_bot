package bot

import (
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

	dg.AddHandler(handlers.InteractionHandler)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Bot started")

	return dg, nil
}
