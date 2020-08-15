package internal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

// Variables used for command line parameters
type Bot struct {
	config    Config
	botPrefix string
	client    *discordgo.Session
	confLoc   string
}

func Start(config Config, configlocation string) {
	fmt.Printf("inside func Start, Prefix set to %s", config.Prefix)
	// Create a new Discord session using the provided bot token.

	client, _ := discordgo.New("Bot " + config.Token)

	bot := Bot{
		confLoc:   configlocation,
		config:    config,
		botPrefix: config.Prefix,
		client:    client,
	}

	// Open a websocket connection to Discord and begin listening.
	if err := client.Open(); err != nil {
		log.Fatalln("Failed to connect to Discord. Is token correct?\n" + err.Error(),
		)
	}

	client.AddHandler(bot.onReady)
	// Register the messageCreate func as a callback for MessageCreate events.
	client.AddHandler(bot.messageCreate)
}
