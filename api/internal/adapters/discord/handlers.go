package discord

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lindseypoche/SELU_ACM/api/internal/adapters/storage"
	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

var (
	// messages
	messageRepository domain.MessageRepository = storage.NewMongoRepo("mongodb://127.0.0.1:27017", "acm", 5)
	messageService    domain.MessageService    = domain.NewMessageService(messageRepository)
)

func validateAuthorAndRole(authorID string, botID string, roles []string) bool {
	if authorID == botID {
		return false
	}
	if roles == nil {
		return false
	}

	return true

	// for i, role := range roles {
	// 	fmt.Printf("%d role: %s\n", i, role)
	// 	// @ACM role id
	// 	if role == "814656414114643969" {
	// 		return true
	// 	}
	// }
	// return false
}

func validateChannel() {

}

// MessageCreated handles messages created
func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {

	if ok := validateAuthorAndRole(m.Author.ID, s.State.User.ID, m.Member.Roles); !ok {
		return
	}

	var attachment *domain.MessageAttachment

	if len(m.Attachments) > 0 {
		attachment = &domain.MessageAttachment{
			ID:       m.Attachments[0].ID,
			URL:      m.Attachments[0].URL,
			Filename: m.Attachments[0].Filename,
			Width:    m.Attachments[0].Width,
			Height:   m.Attachments[0].Height,
			Size:     m.Attachments[0].Size,
		}
	} else {
		attachment = nil
	}

	msg := domain.Message{
		ID:              m.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Content:         m.Content,
		Timestamp:       domain.Timestamp(m.Timestamp),
		EditedTimestamp: domain.Timestamp(m.EditedTimestamp),
		MentionRoles:    m.MentionRoles,
		Attachment:      attachment,
		Author: &domain.User{
			ID:            m.Author.ID,
			Username:      m.Author.Username,
			Discriminator: m.Author.Discriminator,
			Avatar: domain.Avatar{
				ID:       m.Author.Avatar,
				ImageURL: "https:cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".png",
			},
			Email: m.Author.Email,
		},
	}
	b, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	messageService.CreateMessage(msg)
}

// MessageUpdated handles messages updated
func MessageUpdated(s *discordgo.Session, m *discordgo.MessageUpdate) {

	validateAuthorAndRole(m.Author.ID, s.State.User.ID, m.Member.Roles)

	msg := domain.Message{
		ID:              m.Author.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Content:         m.Content,
		Timestamp:       domain.Timestamp(m.Timestamp),
		EditedTimestamp: domain.Timestamp(m.EditedTimestamp),
		MentionRoles:    m.MentionRoles,
		Author: &domain.User{
			ID:            m.Author.ID,
			Username:      m.Author.Username,
			Discriminator: m.Author.Discriminator,
			Avatar: domain.Avatar{
				ID:       m.Author.Avatar,
				ImageURL: "https:cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".png",
			},
		},
	}

	b, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

// MessageReactionAdded handles reactions added to a message
func MessageReactionAdded(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	msgReaction := domain.MessageReaction{
		UserID:    m.UserID,
		MessageID: m.MessageID,
		Emoji: domain.Emoji{
			ID:   m.Emoji.ID,
			Name: m.Emoji.Name,
		},
		ChannelID: m.ChannelID,
	}

	b, err := json.Marshal(&msgReaction)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

// MessageReactionRemoved handles reactions removed
func MessageReactionRemoved(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
}

// GuildMemberAdded handles new guild members
func GuildMemberAdded(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
}

// GuildMemberRemoved handles guild members who were removed or left
func GuildMemberRemoved(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
}

// GuildMemberUpdated handles updated guild members
func GuildMemberUpdated(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
}

// GuildRoleUpdated handles updated roles
func GuildRoleUpdated(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
}

// UserUpdated handles updated user info
func UserUpdated(s *discordgo.Session, m *discordgo.UserUpdate) {
}

// ChannelPinsUpdated : may do something with pins
func ChannelPinsUpdated(s *discordgo.Session, m *discordgo.UserUpdate) {
}
