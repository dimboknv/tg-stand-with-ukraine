package store

type Client struct {
	SentReports map[string]Report
	Session     []byte
}
