package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var (
	Bot *tgbotapi.BotAPI
	err error
)

func InitTelegramBot(wbh, tkn string) {
	Bot, err = tgbotapi.NewBotAPI(tkn)
	if err != nil {
		log.Fatal(err)
	}

	Bot.Debug = false
	log.Printf("Authorized on account %s", Bot.Self.UserName)
	log.Println(wbh + tkn)

	_, err = Bot.SetWebhook(tgbotapi.NewWebhook(wbh + Bot.Token))
	if err != nil {
		log.Fatalln(err)
	}

	info, err := Bot.GetWebhookInfo()
	if err != nil {
		log.Fatalln(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
}
