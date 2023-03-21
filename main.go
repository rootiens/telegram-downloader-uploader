package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rootiens/telegram-downloader-uploader/utils"
)

func main() {
    // Creates temp download directory
	utils.CreateDownloadDir()

    // Database Connection
	utils.ConnectDB()

    
	bot, err := tgbotapi.NewBotAPI(utils.LoadEnv().Token)
	utils.CheckErr(err)

    // Request and Response logs
	bot.Debug = true

    // Create a new UpdateConfig struct with an offset of 0. Offsets are used
    // to make sure Telegram knows we've handled previous values and we don't
    // need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

    // Tell Telegram we should wait up to 30 seconds on each request for an
    // update. This way we can get information just as quickly as making many
    // frequent requests without having to send nearly as many.    
	updateConfig.Timeout = 30

    // Start polling Telegram for updates.
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
