package main

import (
	"TradeBot/commands/fun"
	"TradeBot/commands/info"
	"log"
	"os"

	"TradeBot/bot"
	"TradeBot/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	token = os.Getenv("TOKEN")
	app   = os.Getenv("APP")
	guild = os.Getenv("GUILD")
)

var commandsData = []*discordgo.ApplicationCommand{
	{
		Name:        "echo",
		Description: "Say something",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "message",
				Description: "Say something",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "user_info",
		Description: "Get user info",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "User to inspect",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
		},
	},
}

func main() {
	if token == "" {
		log.Fatal("missing TOKEN")
	}

	dg, err := bot.Start(token)
	if err != nil {
		log.Fatal(err)
	}

	commands.Register("echo", commands.Command{Exec: fun.Echo})
	commands.Register("user_info", commands.Command{Exec: info.UserInfo})

	_, err = dg.ApplicationCommandBulkOverwrite(app, guild, commandsData)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
