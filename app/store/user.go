package store

type User struct {
	Clients map[string]*Client
	Chats   map[int64]*struct {
		Phone      string
		Navigation Navigation
	}
	ID int64
}
