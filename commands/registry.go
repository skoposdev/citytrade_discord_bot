package commands

import "github.com/bwmarrin/discordgo"

type OptionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

type Command struct {
	Exec func(s *discordgo.Session, i *discordgo.InteractionCreate, om OptionMap)
}

var Registry = make(map[string]Command)

func Register(name string, cmd Command) {
	Registry[name] = cmd
}

func Get(name string) (Command, bool) {
	cmd, ok := Registry[name]
	return cmd, ok
}
