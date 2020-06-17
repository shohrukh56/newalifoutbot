package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ReasonCantComeState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var reasonCantComeInstance *ReasonCantComeState

func CreateReasonCantComeStateInstance() *ReasonCantComeState {
	if reasonCantComeInstance == nil {
		reasonCantComeInstance = &ReasonCantComeState{"reasonCantCome", false, false}
	}
	return reasonCantComeInstance
}

func (cantCome *ReasonCantComeState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили!"
}

func (cantCome *ReasonCantComeState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (cantCome *ReasonCantComeState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (cantCome *ReasonCantComeState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (cantCome *ReasonCantComeState) GetName() string {
	return cantCome.name
}

func (cantCome *ReasonCantComeState) IsSendable() bool {
	return cantCome.isSendable
}

func (cantCome *ReasonCantComeState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandHouseholdWork)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandExam)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandTiredWantRest)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandAnotherReason)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (cantCome *ReasonCantComeState) IsTimeRequired() bool {
	return cantCome.isTimeRequired
}
