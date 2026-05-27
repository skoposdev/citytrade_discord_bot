package handlers

import (
	"TradeBot/commands"
	"TradeBot/utils"

	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	name := i.ApplicationCommandData().Name

	cmd, ok := commands.Get(name)
	if !ok {
		return
	}

	opts := utils.ParseOptions(i.ApplicationCommandData().Options)

	cmd.Exec(s, i, opts)
}
