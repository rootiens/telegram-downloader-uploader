package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rootiens/telegram-downloader-uploader/utils"
)

func main() {
	utils.CreateDownloadDir()

	utils.ConnectDB()

	bot, err := tgbotapi.NewBotAPI(utils.LoadEnv().Token)
	utils.CheckErr(err)

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		var textstate string
		chatID := update.Message.Chat.ID
		MessageID := update.Message.MessageID
		UserText := update.Message.Text

		if UserText == "/start" {
			message := "لینک یا فایل مورد نظر خود را ارسال کنید."
			utils.TextMessageSender(bot, message, chatID, MessageID)
		} else if UserText != "" {
			message, method, filename := utils.DownloadFileFromURL(UserText)
			utils.TextMessageSender(bot, message, chatID, MessageID)
			switch method {
			case "upload":
				utils.SendFileByUploading(bot, filename, chatID, MessageID, UserText)
			case "forward":
				utils.SendFileByFileID(bot, filename, chatID, MessageID)
			}
		} else {
			utils.TextMessageSender(bot, textstate, chatID, MessageID)
		}
	}

}
