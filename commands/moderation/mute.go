package moderation

import (
	"TradeBot/commands"
	"TradeBot/commands/moderation/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	ADMIN_ROLE = "1509921612978196662"
)

func Mute(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	u := opts["pseudo"].UserValue(s)
	// t := opts["target"].UserValue(s)

	a, err := s.GuildMember(i.GuildID, i.Member.User.ID)
	if err != nil {
		return
	}

	hasRole := utils.CheckMemberRoles(a, ADMIN_ROLE)

	if !hasRole {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Vous n'avez pas le rôle nécessaire pour effectuer cette action.",
			},
		})
		if err != nil {
			return
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s a bien été mute !", u.DisplayName()),
		},
	})
	if err != nil {
		return
	}
}
