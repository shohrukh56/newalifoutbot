package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	tg "github.com/shohrukh56/newalifoutbot/app/inits"
	. "github.com/shohrukh56/newalifoutbot/app/models"
	"github.com/shohrukh56/newalifoutbot/database"
	"github.com/shohrukh56/newalifoutbot/tools/cache"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AService struct {
}

var SqlArrivalTime sql.NullTime

func InitService() *AService {
	return &AService{}
}

var (
	resCLB    tgbotapi.APIResponse
	err       error
	handlers  = make(map[string]fn)
	actions   = make(map[string]fa)
	SaveState string
)

type fn func(params ...string) tgbotapi.MessageConfig
type fa func(params ...string) tgbotapi.CallbackConfig

func init() {
	handlers[CommandStart] = startHandler
	handlers[Unknown] = unknownMSGHandler
	handlers[CommandBack] = backHandler

	handlers[CommandWillLate] = lateHandler

	handlers[CommandOverslept] = TimeHourHandler
	handlers[CommandTransportProblem] = TimeHourHandler
	handlers[CommandWasInHospital] = TimeHourHandler
	handlers[CommandWasInUniversity] = TimeHourHandler

	handlers[CommandMyChoice] = MyChoiceHandler
	handlers[CommandAnotherReason] = AnotherReasonHandler

	handlers[GetTimeHourInterface] = TimeHourHandler
	handlers[GetTimeMinuteInterface] = TimeMinuteHandler

	handlers[CommandLateButInOffice] = lateButInOfficeHandler

	handlers[CommandMeetingOutOffice] = meetingOutOfficeHandler

	handlers[CommandToPartner] = TimeHourHandler
	handlers[CommandBuyGoodsForOffice] = TimeHourHandler
	handlers[CommandToNationalBank] = TimeHourHandler
	handlers[CommandToTaxInspection] = TimeHourHandler

	handlers[CommandOutOffice] = TimeHourHandler
	handlers[CommandRemoteWork] = TimeHourHandler

	handlers[CommandCantCome] = ReasonCantComeHandler

	handlers[CommandHouseholdWork] = cantComeHandler
	handlers[CommandExam] = cantComeHandler
	handlers[CommandTiredWantRest] = cantComeHandler

	handlers[CommandGetTimeCantComeInterface] = cantComeHandler

	handlers[CommandLastPage] = LastPageHandler

	handlers[CommandOpen] = softUpdateMsgHandler

	actions[ActionUpdate] = softUpdateMsgHandlerCLB
	actions[ActionCheckIn] = checkInHandlerCLB
}

func (s AService) MsgUpdateSrvc() tgbotapi.MessageConfig {
	value, f := cache.Get(strconv.Itoa(update.Message.From.ID))
	if !f {
		user = database.FindUserByID(update.Message.From.ID)
		user.SetState("initial")
		cache.Set(strconv.Itoa(user.UserID), user)
	} else {
		user = value.(User)
	}

	fmt.Printf("#1 User %s set state: %s... \n", user.Username, user.State.GetName())

	prevState := user
	sUserID := strconv.Itoa(user.UserID)
	forwardedMessage := update.Message.Text
	var command string
	state := user.State.GetName()
	fmt.Println("Text forwarded... ", forwardedMessage)
	command = defineCommands(state, forwardedMessage)

	getTimeAndReason(forwardedMessage, state)

	log.Printf("Reason %s, Hour %s, Minute %s, ArrivalTime %s", CommentReason, ArrivalTimeHour, ArrivalTimeMinute, ArrivalTime)

	msg = handlers[command](sUserID, forwardedMessage, state)

	log.Println("Prev state: ", prevState.State.GetName(), "|| forwardedMessage:  ", forwardedMessage, "time req ", prevState.State.IsTimeRequired())

	fmt.Println(ActiveLateButInOffice, "ActiveLateButInOffice")
	fmt.Println(ActiveLate, "ActiveLate")
	fmt.Println(ActiveRemoteWork, "ActiveRemoteWork")
	fmt.Println(ActiveCantCome, "ActiveCantCome")
	fmt.Println(ActiveMeetingOutOfOffice, "ActiveMeetingOutOfOffice")

	if prevState.State.IsSendable() && command != CommandBack {
		if prevState.State.IsTimeRequired() == true && ActiveCantCome == false {
			getTimeAndStore(sUserID, forwardedMessage, prevState)
		} else if command != CommandMyChoice {
			fmt.Println("Time not requered")
			nullT := sql.NullTime{Time: time.Now(), Valid: false}
			MsgID = storeNotification(user, nullT, nullT, prevState.State.GetName())
			fmt.Println(MsgID, "From not required time msgid")
			frm := "Хорошо, спасибо "
			msg.Text = fmt.Sprintf("%s\n %s\n %s\n\nВы ошиблись?\nНажмите, чтобы изменить:\n <a>/open%s</a>\n", CommentReason, ArrivalTime, frm, MsgID)
			msg.ParseMode = "HTML"
			if ENV != "LOCAL" {
				prevState.State.StoreHRDatabase(user, "reason_"+CommentReason+"__"+"absentAt"+ArrivalTime, sql.NullTime{}, MsgID, SaveState)
			}
		}
	}

	return msg
}

func getTimeAndStore(sUserID, forwardedMessage string, prevState User) tgbotapi.MessageConfig {

	SqlArrivalTime.Time, err = getArrivalTime(ArrivalTime)
	SqlArrivalTime.Valid = true
	if err != nil {
		msg = handlers[CommandStart](sUserID, forwardedMessage)
		msg.Text = "<b>😕Упс, видимо вы неправильно указали время или указали время, которое уже прошло</b>\n"
		msg.ParseMode = "HTML"
		return msg
	}
	fmt.Printf("3# Correct format!\n")
	MsgID = storeNotification(user, SqlArrivalTime, sql.NullTime{}, SaveState)
	fmt.Printf("3# Stored notification by %s, %s!\n", user.Username, SqlArrivalTime.Time)
	fmt.Println(MsgID, "From required time msgid")
	frm := "Хорошо, спасибо "
	if ActiveRemoteWork == true {
		frm = "Все приняли, спасибо, удачи😉"
	}
	msg.Text = fmt.Sprintf("%s\n %s\n %s\n\nВы ошиблись?\nНажмите, чтобы изменить:\n <a>/open%s</a>\n", CommentReason, ArrivalTime, frm, MsgID)
	msg.ParseMode = "HTML"
	if ENV != "LOCAL" {
		prevState.State.StoreHRDatabase(user, CommentReason, SqlArrivalTime, MsgID, SaveState)
	}
	return msg
}

func getTimeAndReason(forwardedMessage, state string) {
	if FindReasonCommands(forwardedMessage) == true {
		CommentReason = forwardedMessage
	}
	switch state {
	case "AnotherReason":
		CommentReason = forwardedMessage
	case "MyChoice":
		ArrivalTime = forwardedMessage
	case "TimeHourState":
		ArrivalTimeHour = strings.TrimSuffix(forwardedMessage, "-00")
	case "TimeMinuteState":
		ArrivalTimeMinute = strings.TrimPrefix(forwardedMessage, ArrivalTimeHour+"-")
		ArrivalTime = ArrivalTimeMinute
	}

	if state == "timeCantCome" && forwardedMessage == CommandToday {
		ArrivalTime = time.Now().Format("2006-01-02")
	}
	if state == "timeCantCome" && forwardedMessage == CommandTomorrow {
		date := time.Now().AddDate(0, 0, 1)
		ArrivalTime = date.Format("2006-01-02")
	}
	if forwardedMessage == CommandOutOffice && state == "initial" {
		ActiveOutOfOffice = true
		ActiveCantCome = false
		ActiveRemoteWork = false
		ActiveLate = false
		ActiveLateButInOffice = false
		SaveState = "outOfOffice"
	}

	if forwardedMessage == CommandRemoteWork && state == "initial" {
		ActiveRemoteWork = true
		ActiveCantCome = false
		ActiveLate = false
		ActiveLateButInOffice = false
		ActiveMeetingOutOfOffice = false
		SaveState = "remoteWork"
	}

}

func defineCommands(state, forwardedMessage string) string {
	var command string
	if state == "AnotherReason" && forwardedMessage != "" && forwardedMessage != CommandBack && ActiveCantCome == false {
		command = validateCommand(GetTimeHourInterface)
		return command
	}
	if state == "AnotherReason" && forwardedMessage != "" && forwardedMessage != CommandBack && ActiveCantCome == true {
		command = validateCommand(CommandGetTimeCantComeInterface)
		return command
	}
	if state == "MyChoice" && forwardedMessage != "" && forwardedMessage != CommandBack {
		command = validateCommand(CommandLastPage)
		return command
	}
	if state == "TimeMinuteState" && forwardedMessage != "" && forwardedMessage != CommandBack {
		command = validateCommand(CommandLastPage)
		return command
	}
	if state == "timeCantCome" && (forwardedMessage == CommandToday || forwardedMessage == CommandTomorrow) && ActiveCantCome == true {
		command = validateCommand(CommandLastPage)
		return command
	}

	if state == "TimeHourState" && forwardedMessage != CommandBack && forwardedMessage != CommandMyChoice {
		command = validateCommand(GetTimeMinuteInterface)
		return command
	}
	command = validateCommand(forwardedMessage)
	return command
}

func getArrivalTime(text string) (time.Time, error) {
	var comingTime time.Time
	rule := regexp.MustCompile(`([0-9]|1[0-9]|2[0-3])(\:|\.|.|\-)([0-5][0-9])`)
	timeStr := rule.FindStringSubmatch(text)
	if len(timeStr) == 0 {
		log.Println("Error when find match for regex rule")
		return time.Time{}, errors.New("Wrong time format or omitted")
	}

	replacer := strings.NewReplacer(".", ":", " ", ":", "-", ":")
	timeFormatted := replacer.Replace(timeStr[0])

	date := fmt.Sprintf("%vT%v:00+00:00", time.Now().Format("2006-01-02"), timeFormatted)
	comingTime, err := time.Parse(time.RFC3339, date)

	arrivalTimeUtc, _ := ConvertToUTC(comingTime)

	nowUtc, _ := ConvertToUTC(time.Now())
	nowUtc = nowUtc.Add(time.Hour * 5)

	if arrivalTimeUtc.After(nowUtc) == false && ActiveLateButInOffice == false && ActiveRemoteWork == false {
		log.Println("ERROR when compare arrival time with time now")
		return time.Time{}, errors.New("Wrong time format or omitted")
	}
	return comingTime, err
}

func ConvertToUTC(datetime time.Time) (utcDatetime time.Time, err error) {
	location, _ := time.LoadLocation("UTC")

	utcDatetime = datetime.In(location)

	return utcDatetime, nil
}

func storeNotification(user User, arrivalTime sql.NullTime, leavingTime sql.NullTime, state string) string {

	notification := Notification{
		ChatID:      user.ChatID,
		UserName:    user.Username,
		Status:      state,
		LeavingTime: leavingTime,
		ArrivalTime: arrivalTime,
	}

	msgID, _ := database.CreateNotification(notification)
	return msgID
}

func (s AService) ClbkUpdateSrvc(update tgbotapi.Update) tgbotapi.MessageConfig {
	fmt.Println("Passed callback controller")
	value, f := cache.Get(strconv.Itoa(update.CallbackQuery.From.ID))
	if !f {
		user = database.FindUserByID(update.CallbackQuery.From.ID)
		user.SetState("initial")
		cache.Set(strconv.Itoa(user.UserID), user)
	} else {
		user = value.(User)
	}
	//prevState := user

	action := validateAction(update.CallbackQuery.Data)
	ms := actions[action](update.CallbackQuery.ID, update.CallbackQuery.Data, strconv.Itoa(user.UserID))

	resCLB, err = tg.Bot.AnswerCallbackQuery(ms)
	//
	//msr := tgbotapi.NewMessage(user.ChatID, ms.Text)
	//_, err = tg.Bot.Send(msr)
	//
	//fmt.Println(resMSG.DisableNotification)

	if err != nil {
		log.Errorln(err)
	}

	if resCLB.Ok {
		fmt.Println("*Callback Successful")
	} else {
		fmt.Println("*Callback Issue")
	}
	if err != nil {
		log.Errorln(err)
	}
	return msg
}
