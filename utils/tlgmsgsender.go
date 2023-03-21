package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TextMessageSender(bot *tgbotapi.BotAPI, message string, chatID int64, MessageID int) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Println("\033[31m", err, "\033[0m")
	}
}

func SendFileByUploading(bot *tgbotapi.BotAPI, FileName string, ChatID int64, MessageID int, RequestedURL string) {
	file := tgbotapi.FilePath(FileName)
	msg := tgbotapi.NewDocument(ChatID, file)
	msg.ReplyToMessageID = MessageID
	tlgresp, err := bot.Send(msg)
	if err != nil {
		log.Println("\033[31m", err, "\033[0m")
	}
	res, _ := GetDB().Prepare("insert into files (id, link, fileid, expire_at) values (?,?,?,?)")
	res.Exec(nil, RequestedURL, tlgresp.Document.FileID, nil)
	defer res.Close()
	// log.Println("\033[31m", tlgresp.Document.FileID, "\033[0m")
	go DeleteFileFromDisk(FileName)
}

func SendFileByFileID(bot *tgbotapi.BotAPI, FileID string, ChatID int64, MessageID int) {
	file := tgbotapi.FileID(FileID)
	msg := tgbotapi.NewDocument(ChatID, file)
	msg.ReplyToMessageID = MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Println("\033[31m", err, "\033[0m")
	}
}