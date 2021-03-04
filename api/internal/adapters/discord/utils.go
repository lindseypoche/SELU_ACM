package discord

import "github.com/bwmarrin/discordgo"

// Validate authors
func Validate(s *discordgo.Session, m *discordgo.Message) bool {

	hasAccess := false
	// validate user is not bot
	if m.Author.ID == Config.BotID {
		return hasAccess
	}

	// validate channels
	for _, channel := range Config.Channels {
		if channel == m.ChannelID {
			hasAccess = true
		}
	}

	// validate user roles
	for _, role := range m.Member.Roles {
		for _, accessRole := range Config.Roles {
			if role == accessRole && hasAccess == true {
				return hasAccess
			}
		}
	}
	return hasAccess
}
