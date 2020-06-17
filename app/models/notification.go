package models

import (
	"database/sql"
	"time"
)

type Notification struct {
	MsgID       string       `json:"msg_id"`
	ChatID      int64        `json:"chat_id"`
	UserName    string       `json:"username"`
	LeavingTime sql.NullTime `json:"leaving_time"` //NULLABLE
	ArrivalTime sql.NullTime `json:"arrival_time"`
	IsSent      bool         `json:"is_sent"`
	Status      string       `json:"status"`
	IsDeleted   bool         `json:"is_deleted"`
	CheckedAt   sql.NullTime `json:"checked_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
