package notification

type Notifier interface {
	SendMessage(m *Message) *Result
}

type Message struct {
	To      string
	From    string
	Subject string
	Body    string
}

type Result struct {
	Status  string
	Message string
}
