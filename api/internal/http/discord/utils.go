package discord

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
)

// Validate authors
func Validate(s *discordgo.Session, m *discordgo.Message) bool {

	hasAccess := false
	// validate user is not bot
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
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

func getAttachment(attachments []*discordgo.MessageAttachment) *blogging.MessageAttachment {
	if len(attachments) > 0 {
		return &blogging.MessageAttachment{
			ID:       attachments[0].ID,
			URL:      attachments[0].URL,
			Filename: attachments[0].Filename,
			Width:    attachments[0].Width,
			Height:   attachments[0].Height,
			Size:     attachments[0].Size,
		}
	}
	return nil
}

// snowflakeToUnix converts snowflake id to a unix
func snowflakeToUnix(snowflake string) int {
	v, _ := strconv.Atoi(snowflake)
	x := v>>22 + discordEpoch
	s := strconv.Itoa(x)
	v, _ = strconv.Atoi(s[:len(s)-3])
	return v
}
