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
)

type User struct {
	Clients map[string]*Client  // [phone]client
	Chats   map[int64]*struct { // [chatID]struct
		Phone      string
		Navigation Navigation
	}
	ID int64
}

type Client struct {
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
