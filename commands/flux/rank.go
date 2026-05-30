package flux

import (
	"TradeBot/commands"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func StaffRank(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	u := opts["pseudo"].UserValue(s)
	r := opts["rank"].StringValue()
	chanId := "1486987746391097415"

	embed := &discordgo.MessageEmbed{
		Title:       "Arrivée",
		Color:       0x6dc545,
		Description: fmt.Sprintf("**%s** rejoint l'équipe au sein du pôle **%s** !", u.DisplayName(), r),
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
