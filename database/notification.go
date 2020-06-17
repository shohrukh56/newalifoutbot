package database

import (
	"database/sql"
	"fmt"
	"github.com/shohrukh56/newalifoutbot/app/models"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

func genSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return fmt.Sprintf("%x", id)
}
func CreateNotification(ntf models.Notification) (string, error) {
	ntf.MsgID = genSonyflake()
	var (
		query = "INSERT INTO tg_notifications(msg_id, chat_id, username, leaving_time, arrival_time, is_sent, status) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	)
	pool.Prepare("CreateNotification", query)
	_, err := pool.Exec(query, ntf.MsgID, ntf.ChatID, ntf.UserName, ntf.LeavingTime, ntf.ArrivalTime, ntf.IsSent, ntf.Status)

	if err != nil {
		log.Println("Check CreateNotification: ", err)
		return ntf.MsgID, err
	}

	fmt.Println("New record ID is:", ntf.MsgID)
	return ntf.MsgID, nil
}

func GetNotSentNotifications() ([]models.Notification, error) {
	var (
		query = `SELECT msg_id, username, chat_id, arrival_time, is_sent, status, is_deleted
		FROM tg_notifications
		WHERE is_sent=false AND is_deleted=false AND arrival_time < (now() AT TIME ZONE 'utc-5') AND status = 'late'
		ORDER BY arrival_time;`
		err    error
		ntf    models.Notification
		output []models.Notification
	)
	pool.Prepare("GetNotSentNotifications", query)
	rows, _ := pool.Query(query)

	for rows.Next() {
		err = rows.Scan(&ntf.MsgID, &ntf.UserName, &ntf.ChatID, &ntf.ArrivalTime, &ntf.IsSent, &ntf.Status, &ntf.IsDeleted)
		if err != nil {
			return output, err
		}
		output = append(output, ntf)
	}

	return output, err
}

func UpdateByMsgIDNotification(ids []string) error {
	var (
		err error
	)
	idsStr := strings.Join(ids, ",")
	fmt.Println(idsStr)
	query := fmt.Sprintf("UPDATE tg_notifications SET is_sent=true WHERE msg_id IN (%s)", idsStr)
	fmt.Println(query)
	pool.Prepare("UpdateByMsgIDNotification", query)
	_, err = pool.Exec(query)
	if err != nil {
		log.Println("Check UpdateToIsSentNotification: ", err)
		return err
	}
	return err
}
func SoftDeleteNotification(msgID string) error {
	var (
		err error
	)
	query := fmt.Sprintf("UPDATE tg_notifications SET is_deleted=%v WHERE msg_id = '%s'", true, msgID)
	fmt.Println(query)
	pool.Prepare("SoftDeleteNotification", query)

	_, err = pool.Exec(query)
	if err != nil {
		log.Println("Check SoftDeleteNotification: ", err)
		return err
	}
	return err
}

func UpdateCheckinNotificationByMsgID(msgID string) bool {

	query := fmt.Sprintf("UPDATE tg_notifications SET checked_at=now() WHERE msg_id = '%s' and is_deleted=false and checked_at isnull", msgID)
	fmt.Println(query)
	pool.Prepare("UpdateCheckinNotificationByMsgID", query)
	res, err := pool.Exec(query)
	if err != nil {
		log.Println("Check UpdateCheckinNotificationByMsgID: ", err)
	}
	isAf := true
	if res.RowsAffected() == 0 {
		isAf = false
	}

	return isAf
}

func ExistNotificationByMsgID(msgID string, isCallB bool) (exist bool) {
	t := ""
	if !isCallB {
		t = "is_sent = false and "
	}
	if isCallB {
		t = "checked_at isnull and "
	}
	var query = fmt.Sprintf(`
	SELECT exists (
		SELECT msg_id, is_sent, is_deleted, checked_at 
		FROM tg_notifications
		WHERE msg_id = '%s' and %s is_deleted = false);`, msgID, t)

	fmt.Println("msgID", msgID, " ", query)
	pool.Prepare("ExistNotificationByMsgID", query)
	rows := pool.QueryRow(query)
	err := rows.Scan(&exist)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error checking if row exists '%v' %v", msgID, err)
	}

	log.Println(exist)
	return exist
}
