package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type InitialState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var initialStateInstance *InitialState

func CreateInitialStateInstance() *InitialState {
	if initialStateInstance == nil {
		initialStateInstance = &InitialState{
			"initial",
			false,
			false,
		}
	}
	return initialStateInstance
}

func (initial *InitialState) ResponseOnChangeState() string {
	return "Вы опоздаете или в офисе, но опоздали?"
}

func (initial *InitialState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string, state string) error {
	return nil
}

func (initial *InitialState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (initial *InitialState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (initial *InitialState) GetName() string {
	return initial.name
}

func (initial *InitialState) IsSendable() bool {
	return initial.isSendable
}

func (initial *InitialState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("😬 Опоздаю")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🏢 В офисе, но опоздал")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("👔 По делам офиса")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🚫 Не смогу прийти")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🏃‍♂ По личным делам")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🏠 Работаю из дома")),
	)
}

func (initial *InitialState) IsTimeRequired() bool {
	return initial.isTimeRequired
}
