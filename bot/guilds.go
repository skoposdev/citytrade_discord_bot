package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Guilds = make(map[string]string)

var (
	app = os.Getenv("APP")
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
	{
		Name:        "mute",
		Description: "Permet de muter un membre du discord",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "pseudo",
				Description: "La personne à mute",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
		},
	},
}

func GuildsHandler(s *discordgo.Session) {

	GuildCommands := map[string]map[string]bool{
		"1476574390379614321": {
			"echo":      true,
			"user_info": true,
			"rank":      true,
			"unrank":    true,
			"switch":    true,
			"mute":      true,
		},
		"1482392159863836746": {
			"echo":      true,
			"user_info": true,
		},
	}

	for guildID, cmdRules := range GuildCommands {

		allowedCommands := make([]*discordgo.ApplicationCommand, 0)

		for _, cmd := range commandsData {
			if cmdRules[cmd.Name] {
				allowedCommands = append(allowedCommands, cmd)
			}
		}

		_, err := s.ApplicationCommandBulkOverwrite(app, guildID, allowedCommands)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Déploiement sur la guild: %s\n", guildID)
	}

	for _, guild := range s.State.Guilds {
		g, err := s.Guild(guild.ID)
		if err != nil {
			log.Println("Error getting guild", guild.ID, ":", err)
		}
		Guilds[g.ID] = g.Name
	}
}
