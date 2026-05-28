package handlers

import (
	"fmt"

	"TradeBot/commands"
	"TradeBot/utils"

	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	name := i.ApplicationCommandData().Name
	fmt.Println("interaction:", name)

	cmd, ok := commands.Get(name)
	if !ok {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Unknown command",
			},
		})
		if err != nil {
			fmt.Println("error responding unknown command:", err)
		}
		return
	}

	opts := utils.ParseOptions(i.ApplicationCommandData().Options)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC in command:", name, r)

			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ Command error (crash prevented)",
				},
			})
		}
	}()

	fmt.Println("executing:", name)

	cmd.Exec(s, i, opts)

	fmt.Println("done:", name)
}
