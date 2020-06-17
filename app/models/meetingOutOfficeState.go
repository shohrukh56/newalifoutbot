package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MeetingOutOfficeState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var meetingOutOfficeInstance *MeetingOutOfficeState

func CreateMeetingOutOfficeStateInstance() *MeetingOutOfficeState {
	if meetingOutOfficeInstance == nil {
		meetingOutOfficeInstance = &MeetingOutOfficeState{"meetingOutOffice", false, false}
	}
	return meetingOutOfficeInstance
}

func (meetingOutOffice *MeetingOutOfficeState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили!"
}

func (meetingOutOffice *MeetingOutOfficeState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (meetingOutOffice *MeetingOutOfficeState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (meetingOutOffice *MeetingOutOfficeState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (meetingOutOffice *MeetingOutOfficeState) GetName() string {
	return meetingOutOffice.name
}

func (meetingOutOffice *MeetingOutOfficeState) IsSendable() bool {
	return meetingOutOffice.isSendable
}

func (meetingOutOffice *MeetingOutOfficeState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandToPartner)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBuyGoodsForOffice)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandToNationalBank)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandToTaxInspection)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandAnotherReason)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (meetingOutOffice *MeetingOutOfficeState) IsTimeRequired() bool {
	return meetingOutOffice.isTimeRequired
}
