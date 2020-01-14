package message

type Message struct {
	uid     int
	Message string `json:"message"`
	visible bool
	length  int
	Owner   int `json:"owner"`
}

func (m *Message) Uid() int {
	return m.uid
}
