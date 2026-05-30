package utils

import (
	"github.com/bwmarrin/discordgo"
)

func CheckMemberRoles(member *discordgo.Member, roleTarget string) bool {
	for _, role := range member.Roles {
		if role == roleTarget {
			return true
		}
	}
	return false
}

func CheckRolesPosition(guild *discordgo.Guild, author *discordgo.Member, target *discordgo.Member) bool {
	if target.User.ID == guild.OwnerID {
		return false
	}

	highestAuthorRole, hasAuthorPermission := HighestRolePosition(guild, author)
	highestTargetRole, _ := HighestRolePosition(guild, target)

	return hasAuthorPermission && highestAuthorRole > highestTargetRole
}

func HighestRolePosition(guild *discordgo.Guild, member *discordgo.Member) (int, bool) {
	highestRolePosition := 0
	hasPermission := false

	for _, roleID := range member.Roles {
		for _, role := range guild.Roles {
			if role.ID == roleID {
				if role.Position > highestRolePosition {
					highestRolePosition = role.Position
				}
				if role.Permissions&discordgo.PermissionManageRoles != 0 || role.Permissions&discordgo.PermissionAdministrator != 0 {
					hasPermission = true
				}
				break
			}
		}
	}

	return highestRolePosition, hasPermission
}
