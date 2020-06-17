package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type TimeMinuteState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var timeMinuteStateInstance *TimeMinuteState

func CreateTimeMinuteStateInstance() *TimeMinuteState {
	if timeMinuteStateInstance == nil {
		timeMinuteStateInstance = &TimeMinuteState{"TimeMinuteState", true, true}
	}
	return timeMinuteStateInstance
}

func (late *TimeMinuteState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили об опоздании!"
}

func (late *TimeMinuteState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string, state string) error {
	var (
		client = &http.Client{}
	)
	message := &Res{
		MsgID:       msg_id,
		TGUserID:    user.ID,
		Username:    user.Username,
		Comment:     comment,
		Status:      state,
		ArrivalTime: arrivalTime.Time.Format("2006-01-02 15:04:05"),
	}
	encodedReq, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		// return
	}

	link := defaultstoreLink
	if os.Getenv("employee-time-tracker-logs") != "" {
		link = os.Getenv("employee-time-tracker-logs")
	}
	request, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(encodedReq)))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")

	request.Header.Set("TOKEN", "SOME_RANDOM_STRING")
	if err != nil {
		log.Println(err)
		// return
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		// return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		// return
	}

	log.Println(string(body))

	return nil
}

func (late *TimeMinuteState) SoftUpdateHRDatabase(msg_id string) error {

	var (
		client = &http.Client{}
	)
	message := &Res{
		MsgID: msg_id,
	}
	encodedReq, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		// return
	}

	link := defaultsoftDeleteLink
	if os.Getenv("soft_delete") != "" {
		link = os.Getenv("soft_delete")
	}
	request, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(encodedReq)))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")

	request.Header.Set("TOKEN", "SOME_RANDOM_STRING")
	if err != nil {
		log.Println(err)
		// return
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		// return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		// return
	}

	log.Println(string(body))
	return nil
	return nil
}

func (late *TimeMinuteState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (late *TimeMinuteState) GetName() string {
	return late.name
}

func (late *TimeMinuteState) IsSendable() bool {
	return late.isSendable
}

func (late *TimeMinuteState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	timeHour := strings.TrimSuffix(ArrivalTimeHour, "-00")
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-00"),
			tgbotapi.NewKeyboardButton(timeHour+"-30")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-05"),
			tgbotapi.NewKeyboardButton(timeHour+"-35")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-10"),
			tgbotapi.NewKeyboardButton(timeHour+"-40")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-15"),
			tgbotapi.NewKeyboardButton(timeHour+"-45")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-20"),
			tgbotapi.NewKeyboardButton(timeHour+"-50")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(timeHour+"-25"),
			tgbotapi.NewKeyboardButton(timeHour+"-55")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (late *TimeMinuteState) IsTimeRequired() bool {
	return late.isTimeRequired
}
