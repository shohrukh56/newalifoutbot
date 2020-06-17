package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type LateState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var lateStateInstance *LateState

func CreateLateStateInstance() *LateState {
	if lateStateInstance == nil {
		lateStateInstance = &LateState{"late", false, false}
	}
	return lateStateInstance
}

func (late *LateState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили об опоздании!"
}

func (late *LateState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (late *LateState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (late *LateState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (late *LateState) GetName() string {
	return late.name
}

func (late *LateState) IsSendable() bool {
	return late.isSendable
}

func (late *LateState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandOverslept)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandTransportProblem)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandWasInUniversity)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandWasInHospital)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandAnotherReason)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (late *LateState) IsTimeRequired() bool {
	return late.isTimeRequired
}
