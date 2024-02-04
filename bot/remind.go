package bot

import (
	"bufio"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func saveReminder(message *discordgo.MessageCreate, text string) (bool, string, string, string) {
	reminderRegex := regexp.MustCompile(`^!remind - (.+) - (.+) - (\d{1,2}:\d{2})$`)

	matches := reminderRegex.FindStringSubmatch(text)

	if len(matches) == 0 {
		return false, "", "", ""
	}

	channelID := message.ChannelID
	name := matches[1]
	description := matches[2]
	time := matches[3]

	file, err := os.OpenFile("C:\\Users\\mursa\\OneDrive\\Рабочий стол\\golang project\\discord-bot\\db\\db.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening db.txt:", err)
		return false, "", "", ""
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s - %s - %s - %s\n", channelID, name, description, time)
	if err != nil {
		log.Println("Error writing to db.txt:", err)
		return false, "", "", ""
	}

	return true, name, description, time
}

func listOfReminders(session *discordgo.Session, message *discordgo.MessageCreate) {
	file, err := os.Open("db/db.txt")
	if err != nil {
		fmt.Println("Error opening db.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Пример строки: 1202844171900026882 - Будильник - В 9 утра мне пора в университет, если опоздаю то получу ретейк! - 8:00
		fields := strings.Split(line, " - ")

		if len(fields) != 4 {
			continue
		}

		name := fields[1]
		description := fields[2]
		reminderTime := fields[3]

		reminderTimeParsed, err := time.Parse("15:04", reminderTime)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

		messagee := fmt.Sprintf("Напоминание: %s - %s (время: %s)", name, description, reminderTimeParsed.Format("15:04"))
		session.ChannelMessageSend(message.ChannelID, messagee)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading db.txt:", err)
	}
}
