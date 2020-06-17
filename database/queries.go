package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/shohrukh56/newalifoutbot/app/models"
)

func CreateUser(user models.User) (models.User, error) {
	_, err := pool.Prepare("CreateUser", "INSERT INTO tg_users(username, chat_id, is_active, user_id) VALUES ($1, $2, true, $3) ON CONFLICT ON CONSTRAINT tg_users_pk  DO NOTHING;")
	if err != nil {
		log.Panic(err)
	}
	_, err = pool.Exec("CreateUser", user.Username, user.ChatID, user.UserID)
	if err != nil {
		log.Panic(err)
		return user, err
	}

	return user, nil
}

func FindUserByID(user_id int) (user models.User) {

	_, err := pool.Prepare("FindUserByID", `SELECT id, username, chat_id, user_id, is_active FROM tg_users WHERE user_id = $1;`)
	if err != nil {
		log.Panic(err)
	}
	rows := pool.QueryRow("FindUserByID", user_id)
	err = rows.Scan(
		&user.ID,
		&user.Username,
		&user.ChatID,
		&user.UserID,
		&user.IsActive)
	if err != nil && err != sql.ErrNoRows {
		log.Panic("error checking if row exists '%v' %v", user_id, err)
		return user
	}

	return user
}

func GetCurrentNotifications(chatID int64) ([]models.Notification, error) {
	var (
		query = fmt.Sprintf(`
			SELECT chat_id, msg_id, is_sent, is_deleted, checked_at, leaving_time, arrival_time, status
			FROM tg_notifications
			WHERE chat_id = %d and is_sent = false and is_deleted = false and checked_at isnull;`, chatID)

		err    error
		ntf    models.Notification
		output []models.Notification
	)
	pool.Prepare("GetCurrentNotifications", query)
	fmt.Println(query)
	rows, _ := pool.Query("GetCurrentNotifications")

	for rows.Next() {
		err = rows.Scan(&ntf.ChatID, &ntf.MsgID, &ntf.IsSent, &ntf.IsDeleted, &ntf.CheckedAt, &ntf.LeavingTime, &ntf.ArrivalTime, &ntf.Status)
		if err != nil {
			return output, err
		}
		output = append(output, ntf)
	}

	return output, err
}
