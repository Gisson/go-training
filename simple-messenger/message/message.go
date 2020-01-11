package message

type Message struct {
	uid     int
	message string
	visible bool
	length  int
	owner   int
}

func (m *Message) Owner() int {
	return m.owner
}

func (m *Message) Message() string {
	return m.message
}
