package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := Message{
		ID:              m.Author.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Content:         m.Content,
		Timestamp:       Timestamp(m.Timestamp),
		EditedTimestamp: Timestamp(m.EditedTimestamp),
		MentionRoles:    m.MentionRoles,
		Author: &User{
			m.Author.ID,
			m.Author.Username,
			m.Author.Discriminator,
			m.Author.Avatar,
			m.Author.Email,
		},
	}
	fmt.Println("messageCreate\n\n", msg)
}

func messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := Message{
		ID:              m.Author.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Content:         m.Content,
		Timestamp:       Timestamp(m.Timestamp),
		EditedTimestamp: Timestamp(m.EditedTimestamp),
		MentionRoles:    m.MentionRoles,
		Author: &User{
			m.Author.ID,
			m.Author.Username,
			m.Author.Discriminator,
			m.Author.Avatar,
			m.Author.Email,
		},
	}
	fmt.Println("messageUpdate\n\n", msg)
}

func messageReacted(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.MessageReaction.UserID == s.State.User.ID {
		return
	}

	msgReaction := MessageReaction{
		UserID:    m.UserID,
		MessageID: m.MessageID,
		Emoji:     Emoji{m.Emoji.ID, m.Emoji.Name},
		ChannelID: m.ChannelID,
	}
	fmt.Println("messageReacted:\n\n", msgReaction)
}
