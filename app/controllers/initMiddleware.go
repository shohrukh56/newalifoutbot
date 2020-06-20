package controllers

import (
	"github.com/shohrukh56/newalifoutbot/app/models"
	"github.com/shohrukh56/newalifoutbot/database"
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	user   models.User
	update tgbotapi.Update
)

type ApplicationMiddleWares struct {
	mdl *ApplicationControllers
}

func InitMiddlewares(amdl *ApplicationControllers) *ApplicationMiddleWares {
	return &ApplicationMiddleWares{
		mdl: amdl,
	}
}

func (m *ApplicationMiddleWares) ValidateMessageType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buff := []byte{}
		b := bytes.NewBuffer(buff)
		reader := io.TeeReader(r.Body, b)

		json.NewDecoder(reader).Decode(&update)

		if update.CallbackQuery != nil {
			user.Username = update.CallbackQuery.From.UserName
			user.ChatID = update.CallbackQuery.Message.Chat.ID
			user.UserID = update.CallbackQuery.From.ID
		} else if update.Message != nil {
			user.Username = update.Message.From.UserName
			user.ChatID = update.Message.Chat.ID
			user.UserID = update.Message.From.ID
		} else {
			log.Println("Not provided rule")
		}
		_, err = database.CreateUser(user)
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("#-1 Mdlw update:", update)

		r.Body = ioutil.NopCloser(b)

		next.ServeHTTP(w, r)
	}
}
