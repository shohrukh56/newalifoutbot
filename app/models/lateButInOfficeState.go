package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type LateButInOfficeState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var lateBIOStateInstance *LateButInOfficeState

func CreateLateButInOfficeStateInstance() *LateButInOfficeState {
	if lateBIOStateInstance == nil {
		lateBIOStateInstance = &LateButInOfficeState{"lateButInOffice", false, false}
	}
	return lateBIOStateInstance
}

func (lateBIO *LateButInOfficeState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили об опоздании!"
}

func (lateBIO *LateButInOfficeState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (lateBIO *LateButInOfficeState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (lateBIO *LateButInOfficeState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (lateBIO *LateButInOfficeState) GetName() string {
	return lateBIO.name
}

func (lateBIO *LateButInOfficeState) IsSendable() bool {
	return lateBIO.isSendable
}

func (lateBIO *LateButInOfficeState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
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

func (lateBIO *LateButInOfficeState) IsTimeRequired() bool {
	return lateBIO.isTimeRequired
}
