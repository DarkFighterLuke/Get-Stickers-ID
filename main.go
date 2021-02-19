package main

import (
	"encoding/json"
	"fmt"
	"github.com/NicoNex/echotron"
	"log"
	"os"
	"time"
)

const (
	botLogsFolder = "/Get-Stickers-ID/"
)

type bot struct {
	chatId int64
	echotron.Api
}

var TOKEN = os.Getenv("GetStickersIdBot")
var workingFolder string

func newBot(chatId int64) echotron.Bot {
	return &bot{
		chatId,
		echotron.NewApi(TOKEN),
	}
}

func main() {
	currentPath, _ := os.Getwd()
	workingFolder = currentPath + botLogsFolder

	dsp := echotron.NewDispatcher(TOKEN, newBot)
	dsp.ListenWebhook("https://hiddenfile.tk:443/bot/Get-Stickers-Id", 40988)
}

func (b *bot) Update(update *echotron.Update) {
	b.logUser(update, workingFolder)
	if update.Message.Text == "/start" {
		b.sendStart(update.Message)
	} else {
		b.sendFileID(update.Message)
	}
}

func (b *bot) sendStart(message *echotron.Message) {
	msg := `Hello, %s!
Send me stickers and I will tell you their <b>file_id</b> on Telegram servers.`
	msg = fmt.Sprintf(msg, message.User.FirstName)

	b.SendMessage(msg, message.Chat.ID, echotron.PARSE_HTML)
}

func (b *bot) sendFileID(message *echotron.Message) {
	if message.Sticker == nil {
		b.SendMessage("The message sent is not a sticker.", message.Chat.ID)
	} else {
		fileID := message.Sticker.FileId
		msg := fmt.Sprintf("The sticker <b>file_id</b> is: <code>%s</code>", fileID)
		b.SendMessage(msg, message.Chat.ID, echotron.PARSE_HTML)
	}
}

func (b *bot) logUser(update *echotron.Update, folder string) {
	data, err := json.Marshal(update)
	if err != nil {
		log.Println("Error marshaling logs: ", err)
		return
	}

	os.MkdirAll(folder, 0755)
	var filename string
	if update.Message.Chat.Username == "" {
		filename = folder + update.Message.Chat.FirstName + "_" + update.Message.Chat.LastName + ".txt"
	} else {
		filename = folder + update.Message.Chat.Username + ".txt"
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}

	dataString := time.Now().Format("2006-01-02T15:04:05") + string(data[:])
	_, err = f.WriteString(dataString + "\n")
	if err != nil {
		log.Println(err)
		return
	}
	err = f.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
