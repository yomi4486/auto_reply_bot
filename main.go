package main

import (
	//"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Variables used for command line parameters

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("トークンが読めないんぁああああああああああんでええええええええええ: %v", err)
		os.Exit(1)
	}
	TOKEN := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	myMap := make(map[string]string)
	myMap["hello"] = "こんにちは！"
	myMap["help"] = "Auto-reply botです。"

	if m.Author.ID == s.State.User.ID {
		return
	}

	if value, ok := myMap[m.Content]; ok {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", value))
	}

}
