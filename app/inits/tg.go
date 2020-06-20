package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	Bot *tgbotapi.BotAPI
	err error
	publicURL = os.Getenv("PUBLIC_URL")
	token = os.Getenv("TOKEN")
)

func InitTelegramBot(wbh, tkn string) {
	Bot, err = tgbotapi.NewBotAPI(tkn)
	if err != nil {
		log.Fatal(err)
	}

	Bot.Debug = false
	log.Printf("Authorized on account %s", Bot.Self.UserName)
	log.Println(publicURL+token)

	_, err = Bot.SetWebhook(tgbotapi.NewWebhook(publicURL+token))
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
