package info

import (
	"TradeBot/commands"

	"github.com/bwmarrin/discordgo"
)

func UserInfo(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	u := opts["user"].UserValue(s)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    u.DisplayName(),
			IconURL: u.AvatarURL("16x16"),
		},
		Color: 0xACDDFF,
		Title: "👤  À propos de " + u.Username,
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		return
	}
}
