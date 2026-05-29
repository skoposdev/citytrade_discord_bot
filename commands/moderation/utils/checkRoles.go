package utils

import "github.com/bwmarrin/discordgo"

func CheckMemberRoles(member *discordgo.Member, roleTarget string) bool {
	for _, role := range member.Roles {
		if role == roleTarget {
			return true
		}
	}
	return false
}

func CheckRolesPosition(author *discordgo.Member, target *discordgo.Member, roleTarget string) bool {

}
