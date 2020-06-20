package controllers

import (
	. "github.com/shohrukh56/newalifoutbot/app/models"
	"github.com/shohrukh56/newalifoutbot/database"
	"github.com/shohrukh56/newalifoutbot/tools/cache"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
	"time"
)

var Commands = []string{CommandOverslept, CommandWasInUniversity, CommandWasInHospital, CommandTransportProblem,
	CommandLastPage, CommandMadeMistake, GetTimeMinuteInterface, GetTimeHourInterface,
	CommandStart, CommandBack, CommandWillLate, CommandLateButInOffice, CommandCantCome,
	CommandOutOffice, CommandMeetingOutOffice, CommandRemoteWork, CommandMyChoice, CommandAnotherReason,
	CommandToPartner, CommandBuyGoodsForOffice, CommandToNationalBank, CommandToTaxInspection,
	CommandGetTimeCantComeInterface, CommandHouseholdWork, CommandExam, CommandTiredWantRest, OpenLastNotification, CommandOpen,
}
var ReasonsCommands = []string{CommandWasInHospital, CommandTransportProblem, CommandOverslept,
	CommandWasInUniversity, CommandToPartner, CommandBuyGoodsForOffice, CommandToNationalBank,
	CommandToTaxInspection, CommandHouseholdWork, CommandTiredWantRest, CommandExam, CommandOutOffice, CommandRemoteWork}

var timeCommands = []string{"08-00", "14-00", "09-00", "15-00", "10-00", "16-00", "11-00", "17-00", "12-00", "18-00", "13-00"}

func FindReasonAndTimeCommands(command string, searchCommands []string) bool {
	for _, i := range searchCommands {
		if i == command {
			return true
		}
	}
	return false
}
func validateCommand(command string) string {
	if strings.HasPrefix(command, "/open") {
		command = "/open"
	}
	for _, i := range Commands {
		if i == command {
			return command
		}
	}
	return Unknown
}

func validateAction(action string) string {
	if strings.HasPrefix(action, "/open") {
		return "/open"
	} else if strings.HasPrefix(action, "/checkin") {
		return "/checkin"
	}
	return action
}

func startHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false
	fmt.Println("#2 startHandler() called...")

	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)
	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text = "Здравствуйте! Если опаздываете, опоздали или куда то выходите, напишите мне."
	msg = tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("initial")

	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func unknownMSGHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false

	fmt.Println("#2 unknownMSGHandler() called...")

	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)
	tguID := params[0] //userID
	value, _ := cache.Get(tguID)
	user = value.(User)
	state := params[2]
	text = "😕Упс, неверная команда!\nВыберите, пожалуйста, команду из карточек ниже👇"
	if state == "initial" {
		text = "Здравствуйте! Если опаздываете, опоздали или куда то выходите, напишите мне."
	}
	msg = tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("initial")

	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()
	if user.State.IsSendable() {
		msg.Text = "Спасибо, что сообщили!"
		user.SetState("initial")
		cache.Set(strconv.Itoa(user.UserID), user)
		msg.ReplyMarkup = user.State.GetReplyButtons()
	}

	return msg
}

func lateHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = true
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false
	fmt.Println("#2 lateHandler() called...")
	SaveState = "late"
	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text := "Опаздываете?, Это не есть хорошо, но мы вас понимаем, введите причину опоздания:"
	msg := tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("late")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func backHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 backHandler() called...")
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false
	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	user.SetState("initial")
	cache.Set(strconv.Itoa(user.UserID), user)

	text = "Здравствуйте! Если опаздываете, опоздали или куда то выходите, напишите мне."
	msg = tgbotapi.NewMessage(user.ChatID, text)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func TimeHourHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 TimeHourHandler() called...")

	var user User
	var text string
	fmt.Println(CommentReason)
	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	if ActiveLate == true {
		text = "Примерно во сколько начнете работу? Выберите время часы прихода:"
	}
	if ActiveLateButInOffice == true {
		text = "Во сколько начали работу?"
	}
	if ActiveLateButInOffice == true {
		text = "Во сколько начали работу?"
	}
	if ActiveMeetingOutOfOffice == true || ActiveOutOfOffice == true {
		text = "Примерно во сколько вернетесь? Выберите время часы прихода"
	}
	if ActiveRemoteWork == true {
		text = "Начали работу с дома? Введите, пожалуйста, время часы начало работы "
	}

	msg := tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("TimeHourState")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func AnotherReasonHandler(params ...string) tgbotapi.MessageConfig {

	fmt.Println("#3 AnotherReasonHandler() called...")
	CommentReason = params[1]
	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)
	tguID := params[0] //userID
	value, _ := cache.Get(tguID)
	user = value.(User)

	text = "Введите свой вариант"
	msg = tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("AnotherReason")

	cache.Set(strconv.Itoa(user.UserID), user)
	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func TimeMinuteHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#3 TimeMinuteHandler() called...")
	ArrivalTimeHour = params[1]
	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)
	tguID := params[0] //userID
	value, _ := cache.Get(tguID)
	user = value.(User)
	if ActiveLate == true {
		text = "Примерно во сколько начнете работу? Выберите время минуты прихода:"
	}
	if ActiveLateButInOffice == true {
		text = "Во сколько начали работу?"
	}
	if ActiveMeetingOutOfOffice == true || ActiveOutOfOffice == true {
		text = "Выберите время минуты прихода"
	}
	if ActiveRemoteWork == true {
		text = "Введите, пожалуйста, время минуты начало работы "
	}

	msg = tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("TimeMinuteState")

	cache.Set(strconv.Itoa(user.UserID), user)
	msg.ReplyMarkup = user.State.GetReplyButtons()
	return msg
}

func MyChoiceHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#3 MyChoiceHandler() called...")
	ArrivalTime = params[1]
	var (
		user User
		text string
		msg  tgbotapi.MessageConfig
	)
	tguID := params[0] //userID
	value, _ := cache.Get(tguID)
	user = value.(User)

	text = "Введите свой вариант"
	msg = tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("MyChoice")

	cache.Set(strconv.Itoa(user.UserID), user)
	msg.ReplyMarkup = user.State.GetReplyButtons()
	return msg
}

func lateButInOfficeHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = true
	ActiveMeetingOutOfOffice = false
	fmt.Println("#2 lateButInOfficeHandler() called...")
	SaveState = "lateButInOffice"
	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text := "Опоздали?, Это не есть хорошо, но мы вас понимаем, введите причину опоздания:"
	msg := tgbotapi.NewMessage(user.ChatID, text)
	msg.ParseMode = "HTML"

	user.SetState("lateButInOffice")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func meetingOutOfficeHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = true
	fmt.Println("#2 meetingOutOfficeHandler() called...")
	SaveState = "meetingOutOffice"
	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text := "Вышли с офиса? Выберите причину:"
	msg := tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("meetingOutOffice")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}

func cantComeHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 cantComeHandler() called...")
	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text := "Когда вы не сможете прийти?"
	msg := tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("timeCantCome")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}
func ReasonCantComeHandler(params ...string) tgbotapi.MessageConfig {
	ActiveCantCome = true
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false
	SaveState = "cantCome"
	fmt.Println("#2 ReasonCantComeHandler() called...")

	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	text := "Все в порядке? Почему вы не сможете прийти?"
	msg := tgbotapi.NewMessage(user.ChatID, text)

	user.SetState("reasonCantCome")
	cache.Set(strconv.Itoa(user.UserID), user)

	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}
func clbckCommandBack(params ...string) tgbotapi.CallbackConfig {
	text := ""
	clbID := params[0]
	msg := tgbotapi.NewCallback(clbID, text)
	user.SetState("initial")
	return msg
}
func softUpdateMsgHandlerCLB(params ...string) tgbotapi.CallbackConfig {
	fmt.Println("#2 softUpdateMsgHandlerCLB() called...")

	clbID := params[0]
	prefixText := params[1]

	userID := params[2]

	value, _ := cache.Get(userID)
	user := value.(User)

	msgID := strings.TrimPrefix(prefixText, "/open")

	text := ""
	exist := database.ExistNotificationByMsgID(msgID, true)
	if exist {
		database.SoftDeleteNotification(msgID)
		if ENV != "LOCAL" {
			user.State.SoftUpdateHRDatabase(msgID)
		}
		text = "🆗Запись готова к обновлению. Выберите новую запись"
	}
	text = "Запись не найдена."
	msg := tgbotapi.NewCallback(clbID, text)
	msg.CacheTime = 10
	return msg

}

func LastPageHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 LastPageHandler() called...")
	var user User
	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	user.SetState("initial")
	fmt.Println(MsgID, "MSGID")
	cache.Set(strconv.Itoa(user.UserID), user)
	msg.ReplyMarkup = user.State.GetReplyButtons()

	return msg
}
func softUpdateMsgHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 softUpdateMsgHandler() called...")
	var user User

	tguID := params[0]
	prefixText := params[1]

	value, _ := cache.Get(tguID)
	user = value.(User)
	msgID := strings.TrimPrefix(prefixText, "/open")

	exist := database.ExistNotificationByMsgID(msgID, false)
	msg := startHandler(tguID)
	if exist {
		database.SoftDeleteNotification(msgID)
		if ENV != "LOCAL" {
			user.State.SoftUpdateHRDatabase(msgID)
		}
		msg.Text = "🆗Запись готова к обновлению. Выберите новую запись"
		return msg
	}
	msg.Text = "Запись не найдена."
	return msg
}
func checkInHandlerCLB(params ...string) tgbotapi.CallbackConfig {
	fmt.Println("#2 checkInHandlerCLB() called...")

	clbID := params[0]
	prefixText := params[1]
	userID := params[2]

	value, _ := cache.Get(userID)
	user := value.(User)

	msgID := strings.TrimPrefix(prefixText, "/checkin")
	text := "Запись не найдена."
	nowUtc, _ := ConvertToUTC(time.Now())
	nowUtc = nowUtc.Add(time.Hour * 5)
	if database.UpdateCheckinNotificationByMsgID(msgID) {
		text = "💪Продуктивного дня!"
		if ENV != "LOCAL" {
			user.State.CheckedAtHRDatabase(nowUtc, msgID)
		}

	}
	msg := tgbotapi.NewCallback(clbID, text)
	msg.CacheTime = 10
	return msg

}

func MadeMistakeHandler(params ...string) tgbotapi.MessageConfig {
	fmt.Println("#2 MadeMistakeHandler() called...")
	ActiveCantCome = false
	ActiveRemoteWork = false
	ActiveLate = false
	ActiveLateButInOffice = false
	ActiveMeetingOutOfOffice = false
	var user User

	tguID := params[0]
	value, _ := cache.Get(tguID)
	user = value.(User)

	user.SetState("initial")

	cache.Set(strconv.Itoa(user.UserID), user)
	msg.Text = fmt.Sprintf("Вы хотите изменить запись?")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "/open"+MsgID),
			tgbotapi.NewInlineKeyboardButtonData("Нет", CommandBack)),
	)
	return msg
}

func ActionUpdateHandler(params ...string) tgbotapi.MessageConfig {

	msg.Text = fmt.Sprintf("Вы хотите изменить запись?")
	myButton := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "/open"+MsgID),
			tgbotapi.NewInlineKeyboardButtonData("Нет", CommandBack)),
	)
	msg.ReplyMarkup = myButton
	return msg
}
