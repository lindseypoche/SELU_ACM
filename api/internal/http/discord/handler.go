package discord

import (
	"fmt"
	"log"
	"strings"
	"unsafe"

	"github.com/bwmarrin/discordgo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/architecting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/commenting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/reacting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/mongo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/redis"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/subscribing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/date_utils"
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
	// commenting
	cs = commenting.NewService(new(mongo.CommentRepo))
	// reacting
	rs = reacting.NewService(new(mongo.ReactRepo))
	// subscribing
	ss = subscribing.NewService(new(mongo.SubscribeRepo))
	// architecting
	as = architecting.NewService(new(mongo.ArchitectRepo))
)

// MessageCreated ( current status: ✅ )
// MessageCreated handles messages created
func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {

	// ignore messages from discord pins
	if m.Type == discordgo.MessageTypeChannelPinnedMessage {
		return
	}

	userRoles, getErr := ls.GetUserRoles(m.Author.ID)
	if getErr != nil {
		userRoles = nil
		log.Println("user (bot?) has no saved roles")
	}

	// validation
	accessLvl, ok := Validate(&me{
		authorID:  m.Author.ID,
		stateID:   s.State.User.ID,
		channelID: m.ChannelID,
		isBot:     m.Author.Bot,
		roles:     userRoles,
	})
	// if !ok then bot or unauthorized channel.
	if !ok {
		return
	}

	// return menu
	if m.Content == "!help" {
		menu := helpMenu()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", menu))
		return
	}

	// ignore all < 2 access levels that are not referencing a post.
	if accessLvl < 2 && m.MessageReference == nil {
		return
	}

	// ignore if users are replying to non-event content
	if m.MessageReference != nil {
		_, getErr := ls.GetMessage(m.MessageReference.MessageID)
		if getErr != nil {
			return
		}
	}

	// save all user comments
	if m.MessageReference != nil {
		comment := &commenting.Comment{
			DiscordID: m.ID,
			ChannelID: m.ChannelID,
			Content:   m.Content,
			Timestamp: snowflakeToUnix(m.ID),
			Author: &commenting.User{
				ID:            m.Author.ID,
				Username:      m.Author.Username,
				Discriminator: m.Author.Discriminator,
				Avatar: commenting.Avatar{
					ID:       m.Author.Avatar,
					ImageURL: "https:cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".png",
				},
				Email: m.Author.Email,
			},
			MessageReference: &commenting.MessageReference{
				MessageID: m.MessageReference.MessageID,
				ChannelID: m.MessageReference.ChannelID,
			},
		}

		if err := cs.AddComment(comment); err != nil {
			log.Println("Unable to save comment")
		}
		return
	}

	// officer / admin event creation
	if accessLvl > 1 && m.MessageReference == nil {

		// Validate the created content (start/date, title, body)
		resp, err := parseMessage(m.Content)
		if err != nil {
			// missing content
			// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", prettifyJSON(resp)))
			return
		}

		message := &blogging.Message{
			DiscordID:    m.ID,
			ChannelID:    m.ChannelID,
			GuildID:      m.GuildID,
			StartTime:    resp.Date.Match.i,
			Title:        resp.Title.Match.s,
			Content:      resp.Body.Match.s,
			Timestamp:    snowflakeToUnix(m.ID),
			MentionRoles: m.MentionRoles,
			Attachments:  *(*[]*blogging.MessageAttachment)(unsafe.Pointer(&m.Attachments)),
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

		s.ChannelMessageSend(m.ChannelID, "Post was successfully created")
	}

	// admin access
	// if accessLvl > 2 {

	// 	// save channel
	// 	if strings.Trim(m.Content, " ") == "!acm save channel events" {
	// 		st, err := s.Channel(m.ChannelID)
	// 		if err != nil {
	// 			s.ChannelMessageSend(m.ChannelID, "Unable to retreive channel properties.")
	// 			return
	// 		}

	// 		// Create channel
	// 		restErr := as.CreateChannel(&architecting.Channel{
	// 			DiscordID:  st.ID,
	// 			GuildID:    st.GuildID,
	// 			Name:       st.Name,
	// 			Topic:      st.Topic,
	// 			Collection: "events",
	// 		})
	// 		if restErr != nil {
	// 			s.ChannelMessageSend(m.ChannelID, "Unable to save channel")
	// 			return
	// 		}

	// 		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Channel %s saved successfully.", st.Name))
	// 	}

	// 	if strings.Trim(m.Content, " ") == "!acm save channel officers" {
	// 		st, err := s.Channel(m.ChannelID)
	// 		if err != nil {
	// 			s.ChannelMessageSend(m.ChannelID, "Unable to retreive channel properties.")
	// 			return
	// 		}

	// 		// Create channel
	// 		restErr := as.CreateChannel(&architecting.Channel{
	// 			DiscordID:  st.ID,
	// 			GuildID:    st.GuildID,
	// 			Name:       st.Name,
	// 			Topic:      st.Topic,
	// 			Collection: "officers",
	// 		})
	// 		if restErr != nil {
	// 			s.ChannelMessageSend(m.ChannelID, "Unable to save channel")
	// 			return
	// 		}
	// 		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Channel %s saved successfully.", st.Name))
	// 		return
	// 	}
	// }
}

// MessageUpdated ( current status: ✅ )
func MessageUpdated(s *discordgo.Session, m *discordgo.MessageUpdate) {

	// ignore messages from discord pins
	if m.Type == discordgo.MessageTypeChannelPinnedMessage {
		return
	}

	userRoles, _ := ls.GetUserRoles(m.Author.ID)

	// validation
	accessLvl, ok := Validate(&me{
		authorID:  m.Author.ID,
		stateID:   s.State.User.ID,
		channelID: m.ChannelID,
		isBot:     m.Author.Bot,
		roles:     userRoles,
	})
	if !ok {
		log.Println("not authorized")
		return
	}

	// return menu
	if m.Content == "!help" {
		menu := helpMenu()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", menu))
		return
	}

	// ignore all < 2 access levels that are not referencing a post.
	if accessLvl < 2 && m.MessageReference == nil {
		return
	}

	// ignore if users are replying to non-event content
	if _, getErr := ls.GetMessage(m.MessageReference.MessageID); getErr == nil {
		return
	}

	// anyone can update their comment
	if m.MessageReference != nil {
		comment := &commenting.Comment{
			DiscordID:       m.Message.ID,
			Content:         m.Message.Content,
			EditedTimestamp: 0,
		}
		err := cs.EditComment(comment)
		if err != nil {
			log.Println("error when editing a comment")
		}
		log.Println("comment successfully udpated in db")
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

	// attempt to delete message
	err := bs.DeleteMessage(m.ID, m.ChannelID)
	if err == nil {
		log.Println("Message successfully deleted")
		return
	}

	// attempt to delete comment
	err = cs.DeleteComment(m.ID)
	if err != nil {
		log.Println("Unable to delete comment", err)
		return
	}

	log.Println("Comment successfully deleted")
}

// MessageReactionAdded ( current status: ✅ )
// MessageReactionAdded handles reactions added to a message
func MessageReactionAdded(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

	// no validation, anyone can add an emoji.

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

	// no validation, anyone can remove their emoji

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

// GuildMemberRemoved handles guild members who were removed or left.
func GuildMemberRemoved(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// check user's roles for access
	userRoles, restErr := ls.GetUserRoles(m.User.ID)
	if restErr != nil {
		log.Println(restErr.GetMessage())
		return
	}

	// validation
	if accessLvl, ok := Validate(&me{
		authorID: m.User.ID,
		stateID:  s.State.User.ID,
		isBot:    m.User.Bot,
		roles:    userRoles,
	}); !ok || accessLvl < 1 {
		return
	}

	err := ss.DeleteMember(m.User.ID)
	if err != nil {
		log.Println(err.GetMessage())
	}
}

// BROKEN
// GuildRoleUpdated handles updated events but also creates new roles if
// they don't already exist.
func GuildRoleUpdated(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {

	// ignore roles that dont contain "acm"
	if !strings.Contains(m.GuildRole.Role.Name, "acm") {
		return
	}

	role := architecting.Role{
		DiscordID: m.GuildRole.Role.ID,
		Name:      m.GuildRole.Role.Name,
	}

	err := as.UpdateRole(&role)
	if err != nil {
		log.Println(err.GetMessage())
		return
	}

	log.Println("role successfully updated")
}

func GuildRoleDeleted(s *discordgo.Session, m *discordgo.GuildRoleDelete) {

	err := as.DeleteRole(m.RoleID)
	if err != nil {
		log.Println(err.GetMessage())
		return
	}

	log.Println("role successfully deleted")
}

// ChannelDeleted handles channel deleted events.
func ChannelDeleted(s *discordgo.Session, c *discordgo.ChannelDelete) {}

// ChannelUpdated handles channel update events.
func ChannelUpdated(s *discordgo.Session, c *discordgo.ChannelUpdate) {}

// GuildMemberUpdate updates existing members and saves new members but
// only when a new member's role is updated.
func GuildMemberUpdated(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {

	// Get all acm roles to check against users roles
	acmRoles, err := ls.GetAllRoles()
	if err != nil {
		log.Println(err.GetMessage())
		return
	}

	// Save users acm roles
	var hasAccess bool
	var saveUserRoles []subscribing.Role
	for _, role := range *acmRoles {
		for _, r := range m.Roles {
			if role.DiscordID == r {
				saveUserRoles = append(saveUserRoles, subscribing.Role{
					ID:        role.ID,
					DiscordID: role.DiscordID,
					Name:      role.Name,
				})
				hasAccess = true
			}
		}
	}

	// If user doesn't have access, attempt to delete user
	if !hasAccess {
		err = ss.DeleteMember(m.User.ID)
		if err != nil {
			// couldn't delete member (likely doesn't exist)
			log.Println(err.GetMessage())
			return
		}
		log.Println("member data successfully deleted")
		return
	}

	// Check if user exists in members db, and if they do then update their info
	updateErr := ss.UpdateMember(&subscribing.Member{
		Nick: m.Nick,
		User: &subscribing.User{
			ID:            m.User.ID,
			Username:      m.User.Username,
			Discriminator: m.User.Discriminator,
			Avatar: subscribing.Avatar{
				ID:       m.User.Avatar,
				ImageURL: "https:cdn.discordapp.com/avatars/" + m.User.ID + "/" + m.User.Avatar + ".png",
			},
		},
		Roles: &saveUserRoles,
	})
	if updateErr == nil {
		log.Println("Member was successfully updated")
		return
	} else {
		log.Println(updateErr.GetMessage())
	}

	// If user has acm role but they don't exist in the db, save user.
	saveErr := ss.SaveMember(&subscribing.Member{
		GuildID:  m.GuildID,
		JoinedAt: date_utils.GetNowUnix(),
		Nick:     m.Nick,
		User: &subscribing.User{
			ID:            m.User.ID,
			Username:      m.User.Username,
			Discriminator: m.User.Discriminator,
			Avatar: subscribing.Avatar{
				ID:       m.User.Avatar,
				ImageURL: "https:cdn.discordapp.com/avatars/" + m.User.ID + "/" + m.User.Avatar + ".png",
			},
		},
		Roles: &saveUserRoles,
	})
	if saveErr != nil {
		log.Println(saveErr.GetMessage())
		return
	}

}
