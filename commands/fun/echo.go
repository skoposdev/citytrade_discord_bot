package fun

import (
	"TradeBot/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Echo(s *discordgo.Session, i *discordgo.InteractionCreate, opts commands.OptionMap) {
	builder := new(strings.Builder)

	if v, ok := opts["author"]; ok && v.BoolValue() {
		author := i.Member.User
		builder.WriteString("**" + author.String() + "** says: ")
	}

	builder.WriteString(opts["message"].StringValue())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})
	if err != nil {
		return
	}
}
