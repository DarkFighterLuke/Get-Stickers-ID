package main

import (
	"fmt"
	"github.com/NicoNex/echotron"
	"os"
)

type bot struct {
	chatId int64
	echotron.Api
}

var TOKEN = os.Getenv("GetStickersIdBot")

func newBot(chatId int64) echotron.Bot {
	return &bot{
		chatId,
		echotron.NewApi(TOKEN),
	}
}

func main() {
	dsp := echotron.NewDispatcher(TOKEN, newBot)
	dsp.Poll()
}

func (b *bot) Update(update *echotron.Update) {
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
