package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/reacting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/mongo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/redis"
)

const (
	discordEpoch int = 1420070400000
)

var (
	// implement repositories and services here
	// blogging
	bs = blogging.NewService(new(mongo.BlogRepo), new(redis.BlogCache))
	// listing
	ls = listing.NewService(new(mongo.ListRepo), new(redis.ListCache))
	// reacting
	rs = reacting.NewService(new(mongo.ReactRepo))
)

// MessageCreated ( current status: ✅ )
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
		return
	}

	message := &blogging.Message{
		DiscordID:    m.ID,
		ChannelID:    m.ChannelID,
		GuildID:      m.GuildID,
		Content:      m.Content,
		Timestamp:    snowflakeToUnix(m.ID),
		MentionRoles: m.MentionRoles,
		Attachment:   getAttachment(m.Attachments),
		IsPinned:     m.Pinned,
		Author: &blogging.User{
			ID:            m.Author.ID,
			Username:      m.Author.Username,
			Discriminator: m.Author.Discriminator,
			Avatar: blogging.Avatar{
				ID:       m.Author.Avatar,
				ImageURL: "https:cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".png",
			},
			Email: m.Author.Email,
		},
	}

	if err := bs.AddMessage(message); err != nil {
		log.Println("Unable to save message")
		return
	}
}

// MessageUpdated ( current status: ✅ )
func MessageUpdated(s *discordgo.Session, m *discordgo.MessageUpdate) {

	// TODO: add channel validation (user validation not needed)
	if ok := Validate(s, m.Message); !ok {
		return
	}

	// validate user is not bot
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// get boolean value of the is_pinned property in db
	storedMessage, err := ls.GetMessage(m.ID)
	if err != nil {
		log.Println(err.GetMessage())
		return
	}

	// if message is being pinned, update the original value
	if m.Pinned == true && storedMessage.IsPinned == false {
		err = bs.UpdatePinMessage(m.ID, m.ChannelID, m.Pinned)
		if err != nil {
			log.Println("error when pinning message")
			return
		}
		log.Println("message successfully pinned")
		return
	} else if m.Pinned == false && storedMessage.IsPinned == true {
		// then message is being unpinned. update the original value
		err = bs.UpdatePinMessage(m.ID, m.ChannelID, m.Pinned)
		if err != nil {
			log.Println("error when unpinning message")
			return
		}
		log.Println("message successfully unpinned")
		return
	}

	// if isPinned was not changed, edit the message
	message := &blogging.Message{
		DiscordID:       m.Message.ID,
		Content:         m.Message.Content,
		EditedTimestamp: 0,
	}

	// edit message
	if err := bs.EditMessage(message); err != nil {
		log.Println("error when editing message: ", err)
		return
	}
	log.Println("message successfully edited")
}

// MessageDeleted ( current status: ✅ )
// MessageDeleted handles message deleted reactions
func MessageDeleted(s *discordgo.Session, m *discordgo.MessageDelete) {

	err := bs.DeleteMessage(m.ID, m.ChannelID)
	if err != nil {
		log.Println("Unable to delete message: ", err)
		return
	}
}

// MessageReactionAdded ( current status: ✅ )
// MessageReactionAdded handles reactions added to a message
func MessageReactionAdded(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

	reaction := reacting.MessageReaction{
		UserID:    r.UserID,
		MessageID: r.MessageID,
		ChannelID: r.ChannelID,
		Emoji: reacting.Emoji{
			Name: r.Emoji.Name,
		},
	}

	if err := rs.AddReaction(reaction); err != nil {
		log.Println("Unable to add emoji reaction")
		return
	}
}

// MessageReactionRemoved ( current status: ✅ )
// MessageReactionRemoved handles reactions removed
func MessageReactionRemoved(s *discordgo.Session, r *discordgo.MessageReactionRemove) {

	// reaction := mongo.Emoji{}
	reaction := reacting.MessageReaction{
		UserID:    r.UserID,
		MessageID: r.MessageID,
		Emoji: reacting.Emoji{
			Name: r.Emoji.Name,
		},
	}

	if err := rs.RemoveReaction(reaction); err != nil {
		log.Println("Unable to remove emoji reaction")
		return
	}
}

// https://discord.com/developers/docs/rich-presence/how-to
func PresenceUpdated(s *discordgo.Session, m *discordgo.PresenceUpdate) {
}

func ChannelCreated(s *discordgo.Session, c *discordgo.ChannelCreate) {}

func ChannelUpdated(s *discordgo.Session, c *discordgo.ChannelUpdate) {}

func ChannelDeleted(s *discordgo.Session, c *discordgo.ChannelDelete) {}

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
