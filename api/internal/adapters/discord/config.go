package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	// Bot ...
	Bot   discordBotInterface = &discordSession{}
	BotID string
)

type discordBotInterface interface {
	setBot(*discordgo.Session)
}

type discordSession struct {
	Bot *discordgo.Session
}

type config struct {
	Token     string `json:"token"`
	Prefix    string `json:"prefix"`
	ChannelID string `json:"channel_id"`
	Guild     string `json:"guild"`
}

// Init initializes the bot on start up
func Init() {

	file, err := ioutil.ReadFile("../../internal/clients/discord/config.json")
	if err != nil {
		log.Fatal("Could not read json file: ", err)
	}

	var conf config
	err = json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// set envs
	os.Setenv("BOT_TOKEN", conf.Token)

	fmt.Println("Initializing bot...")
	bot, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		fmt.Println("error making new bot:", err)
		return
	}

	Bot.setBot(bot)

	// Register the messageCreate func as a callback for MessageCreate events.
	bot.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}

func (s *discordSession) setBot(session *discordgo.Session) {
	s.Bot = session
}

func GetSession() discordBotInterface {
	return Bot
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
