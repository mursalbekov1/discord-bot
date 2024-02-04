package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

var BotToken string

func Run() {
	sess, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	sess.AddHandler(newMessage)

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

func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "!info"):
		session.ChannelMessageSend(message.ChannelID, "Привет! Это бот напоминатель! ✏📄\n"+
			"Этот бот умеет ставить напоминалки с различными задачами и временем напоминания. \n\n "+
			"Основные команды: \n"+
			"!remind - 'имя напоминания' : 'текст напоминания' : 'время напоминания' ✅\n"+
			"!remindList - список моих напоминаний 📄\n"+
			"!remindDelete - удалить напоминание по имени -> 'имя' ❌")
	case strings.HasPrefix(message.Content, "!remind"):
		isValid, name, description, remindTime := saveReminder(message, message.Content)
		if isValid {
			session.ChannelMessageSend(message.ChannelID, "Напоминание установлено!")

			reminderTimeParsed, err := time.Parse("15:04", remindTime)
			if err != nil {
				session.ChannelMessageSend(message.ChannelID, "Неверный формат времени.")
				return
			}

			durationUntilReminder := reminderTimeParsed.Sub(time.Now())

			go func() {
				<-time.After(durationUntilReminder)

				session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Напоминание: %s - %s", name, description))
			}()
		} else {
			session.ChannelMessageSend(message.ChannelID, "Неверный формат напоминания.")
		}
	case message.Content == "!list":
		listOfReminders(session, message)
	case strings.HasPrefix(message.Content, "!remindCancel"):
		session.ChannelMessageSend(message.ChannelID, "")
	}
}
