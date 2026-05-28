package main

import (
	"TradeBot/commands/flux"
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
	{
		Name:        "rank",
		Description: "Announce when someone join the team",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "pseudo",
				Description: "La personne ayant été recruté",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "rank",
				Description: "Le pôle de la personne recruté",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "unrank",
		Description: "Announce when someone leave the team",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "pseudo",
				Description: "La personne ayant été unrank",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "switch",
		Description: "Announce when someone switch team",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "pseudo",
				Description: "La personne ayant changé de pôle",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "from",
				Description: "L'ancien pôle du staff",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "to",
				Description: "Le nouveau pôle du staff",
				Type:        discordgo.ApplicationCommandOptionString,
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
	commands.Register("rank", commands.Command{Exec: flux.StaffRank})
	commands.Register("unrank", commands.Command{Exec: flux.StaffUnrank})
	commands.Register("switch", commands.Command{Exec: flux.RankSwitch})

	_, err = dg.ApplicationCommandBulkOverwrite(app, guild, commandsData)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
