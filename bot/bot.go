package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	switch message.Content {
	case "!info":
		session.ChannelMessageSend(message.ChannelID, "Привет! Это бот напоминатель! ✏📄\n"+
			"Этот бот умеет ставить напоминалки с различными задачами и временем напоминания. \n\n "+
			"Основные команды: \n"+
			"!remind - 'имя напоминания' : 'текст напоминания' : 'время напоминания' ✅\n"+
			"!remindList - список моих напоминании 📄\n"+
			"!remindDelete - удалить напоминание по имени -> 'имя' ❌")
	case "!remind":
		session.ChannelMessageSend(message.ChannelID, "Введите имя, текст и время напоминания, разделенные двоеточием. Например: 'Сон:Нужно ложиться спать: 22:00' ✏📄")
	case "!remindList":
		session.ChannelMessageSend(message.ChannelID, "")
	case "!remindCancel":
		session.ChannelMessageSend(message.ChannelID, "")
	}
}
