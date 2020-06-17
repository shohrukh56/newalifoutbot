package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type CantComeState struct {
	name           string
	isSendable     bool
	isTimeRequired bool
}

var cantComeInstance *CantComeState

func CreateCantComeStateInstance() *CantComeState {
	if cantComeInstance == nil {
		cantComeInstance = &CantComeState{"timeCantCome", true, false}
	}
	return cantComeInstance
}

func (cantCome *CantComeState) ResponseOnChangeState() string {
	return "Спасибо, что уведомили!"
}

func (cantCome *CantComeState) StoreHRDatabase(user User, comment string, arrivalTime sql.NullTime, msg_id string,state string) error {
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

func (cantCome *CantComeState) SoftUpdateHRDatabase(msg_id string) error {

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
}

func (cantCome *CantComeState) CheckedAtHRDatabase(chk time.Time, msg_id string) error {
	return nil
}

func (cantCome *CantComeState) GetName() string {
	return cantCome.name
}

func (cantCome *CantComeState) IsSendable() bool {
	return cantCome.isSendable
}

func (cantCome *CantComeState) GetReplyButtons() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandToday)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandTomorrow)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandMyChoice)),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandBack)),
	)
}

func (cantCome *CantComeState) IsTimeRequired() bool {
	return cantCome.isTimeRequired
}
