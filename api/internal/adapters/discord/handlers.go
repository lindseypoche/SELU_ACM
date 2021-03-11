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
	// _, _ = s.ChannelMessageSend(m.ChannelID, "access denied: your message was not posted to the web")

	msg := domain.Message{
		ID:           m.ID,
		ChannelID:    m.ChannelID,
		GuildID:      m.GuildID,
		Content:      m.Content,
		Timestamp:    snowflakeToUnix(m.ID),
		MentionRoles: m.MentionRoles,
		Attachment:   getAttachment(m.Attachments),
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

	resp, err := messageService.CreateMessage(msg)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, resp.Success)
}

// should be changed to MessageEdited
// MessageUpdated handles messages updated (WORKING)
func MessageUpdated(s *discordgo.Session, m *discordgo.MessageUpdate) {

	if ok := Validate(s, m.Message); !ok {
		// _, _ = s.ChannelMessageSend(m.ChannelID, "access denied: your edit was not updated on the web")
		return
	}

	// editedTimestamp := date_utils.GetNowUnix()
	// fmt.Println(editedTimestamp)

	fmt.Println("m.Pinned: ", m.Pinned)
	if m.Pinned == true {
		// update LatestPin in channel
		pin := &domain.Pin{
			ChannelID: m.ChannelID,
			Message: &domain.Message{
				ID:         m.ID,
				ChannelID:  m.ChannelID,
				Content:    m.Content,
				Pinned:     m.Pinned,
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
			_, _ = s.ChannelMessageSend(m.Message.ChannelID, err.GetMessage())
			return
		}
		_, _ = s.ChannelMessageSend(m.Message.ChannelID, fmt.Sprintf("channel_id: %s, did_pin: %v", pin.ChannelID, m.Pinned))
		return
	}

	msg := &domain.Message{
		ID:              m.Message.ID,
		Content:         m.Message.Content,
		EditedTimestamp: 0,
	}

	resp, restErr := messageService.UpdateMessage(msg)
	if restErr != nil {
		_, _ = s.ChannelMessageSend(m.Message.ChannelID, restErr.GetMessage())
		return
	}
	_, _ = s.ChannelMessageSend(m.Message.ChannelID, fmt.Sprintf("edited_timestamp: %d", msg.EditedTimestamp))
	_, _ = s.ChannelMessageSend(m.Message.ChannelID, resp.Success)
}

// MessageDeleted handles message deleted reactions
func MessageDeleted(s *discordgo.Session, m *discordgo.MessageDelete) {

	id := m.ID
	restErr := messageService.DeleteMessage(id)
	if restErr != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, restErr.GetMessage())
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("message %s successfully deleted", id))
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
		_, _ = s.ChannelMessageSend(r.ChannelID, restErr.GetMessage())
		return
	}
	_, _ = s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("emoji %s added to db", r.Emoji.Name))
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
		_, _ = s.ChannelMessageSend(r.ChannelID, restErr.GetMessage())
		return
	}
	_, _ = s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("emoji %s deleted from db", r.Emoji.Name))
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
