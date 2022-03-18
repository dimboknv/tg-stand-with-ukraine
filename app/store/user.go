package store

import "time"

type Navigation int

const (
	LoginCommand  = "login"
	LogoutCommand = "logout"
	StartCommand  = "start"

	WelcomeNavigation Navigation = iota
	PhoneNavigation
	CodeNavigation
	Pass2faNavigation
	UserNavigation
	SharePhoneNavigation
	SplitCode1Navigation
	SplitCode2Navigation
)

type User struct {
	Clients map[string]*Client // [phone]client
	Chats   map[int64]*Chat    // [chatID]struct
	Phone   string
	ID      int64
}

type Chat struct {
	AuthPhone    string
	AuthCode     string
	DeleteMsgIDs []int
	ID           int64
	ReplyMsgID   int
	Navigation   Navigation
}

type Client struct {
	LastConnectionAt time.Time
	SignInAt         time.Time
	SentReports      map[string]Report // [url]report
	Phone            string
	Session          []byte
	IsAuthorized     bool
}

type Report struct {
	CreatedAt time.Time
	Text      string
	URL       string
}

type Rashist struct {
	CreatedAt  time.Time
	URL        string
	ResolveErr string
}
