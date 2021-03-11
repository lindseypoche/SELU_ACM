package discord

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/lindseypoche/SELU_ACM/api/internal/adapters/storage"
	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

var (
	// messages
	messageRepository domain.MessageRepository = storage.NewMongoRepo("mongodb://127.0.0.1:27017", "acm", 5)
	messageService    domain.MessageService    = domain.NewMessageService(messageRepository)
)

const (
	discordEpoch int = 1420070400000
)

func getAttachment(attachments []*discordgo.MessageAttachment) *domain.MessageAttachment {
	if len(attachments) > 0 {
		return &domain.MessageAttachment{
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

// MessageCreated handles messages created (WORKING)
func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {

	if ok := Validate(s, m.Message); !ok {
		return
	}

	// validate user is not bot
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// ignore messages from discord pins
	if m.Type == discordgo.MessageTypeChannelPinnedMessage {
		fmt.Println("pinned messages from bot are not read")
		return
	}

	msg := domain.Message{
		ID:           m.ID,
		ChannelID:    m.ChannelID,
		GuildID:      m.GuildID,
		Content:      m.Content,
		Timestamp:    snowflakeToUnix(m.ID),
		MentionRoles: m.MentionRoles,
		Attachment:   getAttachment(m.Attachments),
		Pinned:       m.Pinned,
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

	_, err := messageService.CreateMessage(msg)
	if err != nil {
		return
	}
}

// should be changed to MessageEdited
// MessageUpdated handles messages updated (WORKING)
func MessageUpdated(s *discordgo.Session, m *discordgo.MessageUpdate) {

	if ok := Validate(s, m.Message); !ok {
		return
	}

	// validate user is not bot
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	fmt.Println("m.Pinned: ", m.Pinned)
	if m.Pinned == true {
		// update LatestPin in channel
		pin := &domain.Pin{
			MessageID: m.ID,
			ChannelID: m.ChannelID,
			Message: &domain.Message{
				ID:         m.ID,
				ChannelID:  m.ChannelID,
				Content:    m.Content,
				Pinned:     true,
				Attachment: getAttachment(m.Attachments),
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
			},
			PinnedAt: 0,
		}
		err := messageService.UpdateLatestPin(pin)
		if err != nil {
			return
		}
		return
	}

	msg := &domain.Message{
		ID:              m.Message.ID,
		Content:         m.Message.Content,
		EditedTimestamp: 0,
	}

	_, restErr := messageService.UpdateMessage(msg)
	if restErr != nil {
		return
	}
}

// MessageDeleted handles message deleted reactions
func MessageDeleted(s *discordgo.Session, m *discordgo.MessageDelete) {

	message := domain.Message{
		ID:        m.ID,
		ChannelID: m.ChannelID,
		Pinned:    m.Pinned,
	}
	restErr := messageService.DeleteMessage(&message)
	if restErr != nil {
		return
	}
}

// MessageReactionAdded handles reactions added to a message
func MessageReactionAdded(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

	reaction := domain.MessageReaction{
		UserID:    r.UserID,
		MessageID: r.MessageID,
		Emoji: domain.Emoji{
			Name: r.Emoji.Name,
		},
	}

	restErr := messageService.UpdateReaction(reaction)
	if restErr != nil {
		return
	}
}

// MessageReactionRemoved handles reactions removed
func MessageReactionRemoved(s *discordgo.Session, r *discordgo.MessageReactionRemove) {

	// reaction := domain.Emoji{}
	reaction := domain.MessageReaction{
		UserID:    r.UserID,
		MessageID: r.MessageID,
		Emoji: domain.Emoji{
			Name: r.Emoji.Name,
		},
	}

	restErr := messageService.RemoveReaction(reaction)
	if restErr != nil {
		return
	}
}

// ChannelPinsUpdated : may do something with pins
func ChannelPinsUpdated(s *discordgo.Session, m *discordgo.ChannelPinsUpdate) {

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
