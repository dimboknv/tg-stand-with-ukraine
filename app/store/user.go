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
)

type User struct {
	Clients map[string]*Client // [phone]client
	Chats   map[int64]*Chat    // [chatID]struct
	Phone   string
	ID      int64
}

type Chat struct {
	AuthPhone         string
	ID                int64
	ShareContactMsgID int
	Navigation        Navigation
}

type Client struct {
	Phone       string
	SentReports map[string]Report // [url]report
	Session     []byte
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
