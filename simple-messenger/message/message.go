package message

type Message struct {
	uid     int
	message string `json:"message"`
	visible bool
	length  int
	owner   int `json:"owner"`
}

func (m *Message) Owner() int {
	return m.owner
}

func (m *Message) Message() string {
	return m.message
}
