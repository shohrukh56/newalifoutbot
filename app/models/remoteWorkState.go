package models

import (
	"database/sql"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type RemoteWorkState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var remoteWorkInstance *RemoteWorkState

func CreateRemoteWorkStateInstance() *RemoteWorkState {
	if remoteWorkInstance == nil {
		remoteWorkInstance = &RemoteWorkState{"remoteWork", true, false}
	}
	return remoteWorkInstance
}

func (remoteWork *RemoteWorkState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили!"
}

func (remoteWork *RemoteWorkState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
	return nil
}

func (remoteWork *RemoteWorkState) SoftUpdateHRDatabase(msg_id string) error {
	return nil
}

func (remoteWork *RemoteWorkState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (remoteWork *RemoteWorkState) GetName() string {
	return remoteWork.name
}

func (remoteWork *RemoteWorkState) IsSendable() bool {
	return remoteWork.isSendable
}

func (remoteWork *RemoteWorkState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⬅ Вернуться"),
		),
	)
}

func (remoteWork *RemoteWorkState) IsTimeRequired() bool {
	return remoteWork.isTimeRequired
}
