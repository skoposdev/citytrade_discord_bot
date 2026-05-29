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
	guild, _ := s.Guild(i.GuildID)

	p := opts["pseudo"].UserValue(s)
	// t := opts["target"].UserValue(s)

	author, err := s.GuildMember(i.GuildID, i.Member.User.ID)
	if err != nil {
		return
	}

	target, err := s.GuildMember(i.GuildID, p.ID)
	if err != nil {
		return
	}

	canMute := utils.CheckRolesPosition(guild, author, target)

	hasRole := utils.CheckMemberRoles(author, ADMIN_ROLE)

	if !hasRole || !canMute {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Vous n'avez pas le rôle ou les permissions nécessaires nécessaire pour effectuer cette action.",
			},
		})
		if err != nil {
			return
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s a bien été mute !", target.DisplayName()),
		},
	})
	if err != nil {
		return
	}
}
