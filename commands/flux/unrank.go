package flux

import (
	"TradeBot/commands"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func StaffUnrank(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	u := opts["pseudo"].StringValue()
	chanId := "1486987746391097415"

	embed := &discordgo.MessageEmbed{
		Title:       "Départ",
		Color:       0xe06666,
		Description: fmt.Sprintf("**%s** quitte malheureusement notre équipe...", u),
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
