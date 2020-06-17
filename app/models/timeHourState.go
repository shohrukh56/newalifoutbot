
package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TimeHourState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var timeHourStateInstance *TimeHourState

func CreateTimeHourStateInstance() *TimeHourState {
	if timeHourStateInstance == nil {
		timeHourStateInstance = &TimeHourState{"TimeHourState", false, false}
	}
	return timeHourStateInstance
}

func (late *TimeHourState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили об опоздании!"
}

func (late *TimeHourState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (late *TimeHourState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (late *TimeHourState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (late *TimeHourState) GetName() string {
	return late.name
}

func (late *TimeHourState) IsSendable() bool {
	return late.isSendable
}

func (late *TimeHourState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("08-00"),
			tgbotapi.NewKeyboardButton("14-00")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("09-00"),
			tgbotapi.NewKeyboardButton("15-00")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("10-00"),
			tgbotapi.NewKeyboardButton("16-00")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("11-00"),
			tgbotapi.NewKeyboardButton("17-00")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("12-00"),
			tgbotapi.NewKeyboardButton("18-00")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("13-00"),
			tgbotapi.NewKeyboardButton(CommandMyChoice)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (late *TimeHourState) IsTimeRequired() bool {
	return late.isTimeRequired
}

