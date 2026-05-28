package flux

import (
	"TradeBot/commands"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RankSwitch(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	u := opts["pseudo"].UserValue(s)
	fromRank := opts["from"].StringValue()
	toRank := opts["to"].StringValue()
	chanId := "1486987746391097415"

	embed := &discordgo.MessageEmbed{
		Title:       "Changement de pôle",
		Color:       0xf1c232,
		Description: fmt.Sprintf("**%s** quitte son poste de **%s** et devient **%s** !", u.DisplayName(), fromRank, toRank),
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Flux envoyé ✅",
		},
	})
	if err != nil {
		fmt.Println("interaction error:", err)
		return
	}

	_, err = s.ChannelMessageSendComplex(chanId, &discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{embed}})

	if err != nil {
		return
	}
}
