package models

const (
	ActionCheckIn          = "/checkin"
	ActionUpdate            = "/open"
	CommandStart            = "/start"
	CommandBack             = "⬅ Вернуться"
	Unknown                 = "unknown"
	CommandMadeMistake      = "Ошибся"
	CommandLastPage         = "Последнее окно"
	ENV                     = "LOCAL"
	OpenLastNotification    = "/OpenLastNotification"
	CommandOpen             = "/open"
	CommandWillLate         = "😬 Опоздаю"
	CommandLateButInOffice  = "🏢 В офисе, но опоздал"
	CommandCantCome         = "🚫 Не смогу прийти"
	CommandOutOffice        = "🏃‍♂ По личным делам"
	CommandMeetingOutOffice = "👔 По делам офиса"
	CommandRemoteWork       = "🏠 Работаю из дома"

	CommandOverslept        = "Проспал"
	CommandTransportProblem = "Транспортные проблемы"
	CommandWasInUniversity  = "В универе был"
	CommandWasInHospital    = "Был у врача"

	CommandMyChoice      = "Свой вариант"
	CommandAnotherReason = "Другая причина"

	GetTimeHourInterface   = "Опция время часы"
	GetTimeMinuteInterface = "Опция время минуты"

	CommandToPartner         = "Еду к партнеру"
	CommandBuyGoodsForOffice = "Еду купить товары для офиса"
	CommandToNationalBank    = "Нужно в Нац банк"
	CommandToTaxInspection   = "Нужно в налоговую"

	CommandToday    = "Сегодня"
	CommandTomorrow = "Завтра"

	CommandHouseholdWork            = "По домашним делам"
	CommandExam                     = "Экзамен"
	CommandTiredWantRest            = "Устал, чуть отдохну"
	CommandGetTimeCantComeInterface = "Дай время несмогу придти"
)

var (
	CommentReason     string
	ArrivalTimeHour   string
	ArrivalTimeMinute string
	ArrivalTime       string
	MsgID             string
)
var (
	ActiveOutOfOffice        bool
	ActiveMeetingOutOfOffice bool
	ActiveLate               bool
	ActiveLateButInOffice    bool
	ActiveRemoteWork         bool
	ActiveCantCome           bool
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	ChatID   int64  `json:"chat_id"`
	UserID   int    `json:"user_id"`
	State    State  `json:"state"`
	IsActive bool   `json:"is_active"`
}

func (uc *User) SetState(state string) {
	switch state {
	case "late":
		uc.State = CreateLateStateInstance()
	case "lateButInOffice":
		uc.State = CreateLateButInOfficeStateInstance()
	case "initial":
		uc.State = CreateInitialStateInstance()
	case "timeCantCome":
		uc.State = CreateCantComeStateInstance()
	case "reasonCantCome":
		uc.State = CreateReasonCantComeStateInstance()
	case "meetingOutOffice":
		uc.State = CreateMeetingOutOfficeStateInstance()
	case "remoteWork":
		uc.State = CreateRemoteWorkStateInstance()
	case "TimeHourState":
		uc.State = CreateTimeHourStateInstance()
	case "TimeMinuteState":
		uc.State = CreateTimeMinuteStateInstance()
	case "AnotherReason":
		uc.State = CreateAnotherReasonStateInstance()
	case "MyChoice":
		uc.State = CreateMyChoiceStateInstance()
	default:
		uc.State = CreateInitialStateInstance()
	}

}
