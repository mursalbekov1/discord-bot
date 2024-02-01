package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sess, err := discordgo.New("Bot MTIwMjUxNTM5NDc3NTA4OTI0Mg.GJvSxW.2Pik1nvPtWCPTTvDoFaloSazmLXFXBq9pKR1NE")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "Hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		} else {
			s.ChannelMessageSend(m.ChannelID, "The command is not correct!")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is working!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
