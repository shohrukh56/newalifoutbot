package controllers

import (
	tg "github.com/shohrukh56/newalifoutbot/app/inits"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ApplicationControllers struct {
	srv *AService
}

func InitControllers(asrv *AService) *ApplicationControllers {
	return &ApplicationControllers{
		srv: asrv,
	}
}
var (
	msg    tgbotapi.MessageConfig
)
func (c *ApplicationControllers) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("## Passed webHookHandler...")

		defer r.Body.Close()

		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		err = json.Unmarshal(bytes, &update)

		if err !=nil{
			log.Println(err)
		}

		if update.CallbackQuery != nil {
			msg = c.srv.ClbkUpdateSrvc(update)
		}

		if update.Message != nil {
			msg = c.srv.MsgUpdateSrvc()
		}



		fmt.Println(user,"user")
		fmt.Println("Check before sending: ", msg.Text, msg.ChatID)

		_, err = tg.Bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

//**************************
