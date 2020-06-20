package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shohrukh56/newalifoutbot/app/controllers"
	"github.com/shohrukh56/newalifoutbot/app/inits"
	"github.com/shohrukh56/newalifoutbot/configs"
	"github.com/shohrukh56/newalifoutbot/database"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	config = configs.TGConfig()
	wbh    = config.WebHookURL
	tkn    = config.Token
)

func main() {

	initLogging()
	tg.InitTelegramBot(wbh, tkn)
	log.Println("bot initialized")
	database.Connect(configs.DBConfig())

	appSrvc := controllers.InitService()

	appContrl := controllers.InitControllers(appSrvc)
	appMiddle := controllers.InitMiddlewares(appContrl)

	handler := appContrl.WebhookHandler()

	handleNull := appMiddle.ValidateMessageType(handler)

	mux := http.NewServeMux()
	mux.HandleFunc("/"+config.Token, handleNull)
	go periodicalWorker(time.Minute, sendMessage)

	http.ListenAndServe("0.0.0.0:80", mux)

}

func sendMessage() {

	var ids []string
	notifications, err := database.GetNotSentNotifications()
	fmt.Println(notifications, "notifications")
	if err != nil {
		log.Println(err)
	}

	if len(notifications) > 0 {
		for _, notification := range notifications {
			log.Infof("## sendMessage() Notification [%s]...", notification.UserName)
			ids = append(ids, "'"+notification.MsgID+"'")
			msgID := notification.MsgID
			msg := tgbotapi.NewMessage(notification.ChatID, "<b>Уведомление</b>\n\nВы в офисе?")
			msg.ParseMode = "HTML"
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("В офисе", "/checkin"+msgID),
					tgbotapi.NewInlineKeyboardButtonData("Изменить запись", "/open"+msgID),
				),
			)
			tg.Bot.Send(msg)
		}
		database.UpdateByMsgIDNotification(ids)
	}

}
func periodicalWorker(d time.Duration, f func()) {

	var reentranceFlag int64

	for range time.Tick(d) {

		go func() {
			if atomic.CompareAndSwapInt64(&reentranceFlag, 0, 1) {
				defer atomic.StoreInt64(&reentranceFlag, 0)
			} else {
				log.Println("Previous worker in process now")
				return
			}
			f()
		}()

	}
}

func initLogging() {

	file, err := os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{})
}
