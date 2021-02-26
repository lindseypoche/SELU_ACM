package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"

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

	for i, role := range roles {
		fmt.Printf("%d role: %s\n", i, role)
		// @ACM role id
		if role == "814656414114643969" {
			return true
		}
	}
	return false
}

func serveFrames(imgByte []byte) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./img.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}

func validateChannel() {

}

// MessageCreated handles messages created
func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {

	if ok := validateAuthorAndRole(m.Author.ID, s.State.User.ID, m.Member.Roles); !ok {
		return
	}

	msg := domain.Message{
		ID:              m.ID,
		ChannelID:       m.ChannelID,
		GuildID:         m.GuildID,
		Content:         m.Content,
		Timestamp:       domain.Timestamp(m.Timestamp),
		EditedTimestamp: domain.Timestamp(m.EditedTimestamp),
		MentionRoles:    m.MentionRoles,
		// Attachments: []*domain.MessageAttachment{
		// only get first attachment. should replace with for loop function
		// &domain.MessageAttachment{
		// 	ID:       m.Attachments[0].ID,
		// 	URL:      m.Attachments[0].URL,
		// 	Filename: m.Attachments[0].Filename,
		// 	Width:    m.Attachments[0].Width,
		// 	Height:   m.Attachments[0].Height,
		// 	Size:     m.Attachments[0].Size,
		// },
		// },
		Author: &domain.User{
			ID:            m.Author.ID,
			Username:      m.Author.Username,
			Discriminator: m.Author.Discriminator,
			Avatar:        m.Author.Avatar,
			Email:         m.Author.Email,
		},
	}

	// TODO: decode image and store on server. store path in db.
	// img, err := s.UserAvatarDecode(m.Member.User)
	// img, err := s.UserAvatarDecode(s.State.User)
	// if err != nil {
	// 	_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
	// 	return
	// }
	// image.DecodeConfig(image.DecodeConfig(io.Reader))
	// fmt.Println("img:>>> ", img)

	// result, err := domain.MessageSrvc.CreateMessage(msg)
	result, err := messageService.CreateMessage(msg)
	// b, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, result.Success)
	// fmt.Println(string(b))
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
			Avatar:        m.Author.Avatar,
			Email:         m.Author.Email,
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
