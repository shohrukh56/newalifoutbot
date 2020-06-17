package models

import (
	"database/sql"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	defaultstoreLink      = "https://hr.alif.tj/api/employee-time-tracker-logs"
	defaultupdateLink     = "https://hr.alif.tj/api/employee-time-tracker-logs-update"
	defaultsoftDeleteLink = "https://hr.alif.tj/api/soft_delete"
	storeLink             = os.Getenv("LogsStore")
	updateLink            = os.Getenv("LogsUpdate")
	softDeleteLink        = os.Getenv("LogsSoftUpdate")
)

type Res struct {
	MsgID       string `json:"msg_id"`
	TGUserID    int    `json:"tg_user_id"`
	Username    string `json:"userName"`
	Comment     string `json:"comment"`
	Status      string `json:"status"`
	LeavingTime string `json:"leaving_time,omitempty"`
	ArrivalTime string `json:"arrival_time,omitempty"`
	CheckedAt   string `json:"checked_at,omitempty"`
}

//2.Реализация интерфейса States(действия выполняемые при конкретном Состоянии)
type State interface {
	ResponseOnChangeState() string
	StoreHRDatabase(User, string, sql.NullTime, string,string) error
	SoftUpdateHRDatabase(string) error
	CheckedAtHRDatabase(time.Time, string) error
	GetName() string
	GetReplyButtons() tgbotapi.ReplyKeyboardMarkup
	IsSendable() bool
	IsTimeRequired() bool
}
