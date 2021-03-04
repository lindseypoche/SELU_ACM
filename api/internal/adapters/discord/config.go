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
	Token    string            `json:"token"`
	Owners   []string          `json:"owners"`
	Prefix   string            `json:"prefix"`
	Channels map[string]string `json:"channels"`
	Guild    string            `json:"guild"`
}

// Init initializes the bot on start up
func Init() {

	file, err := ioutil.ReadFile("../../internal/adapters/discord/cake.json")
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

	// Register handlers
	bot.AddHandler(MessageCreated)
	bot.AddHandler(MessageUpdated)
	bot.AddHandler(MessageReactionAdded)
	bot.AddHandler(MessageReactionRemoved)
	bot.AddHandler(GuildMemberAdded)
	bot.AddHandler(GuildMemberRemoved)
	bot.AddHandler(GuildMemberUpdated)
	bot.AddHandler(UserUpdated)

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
