package main

import (
	_ "bufio"
	"errors"
	_ "fmt"
	"github.com/Gisson/simple-messenger/api"
	msgpkg "github.com/Gisson/simple-messenger/message"
	"github.com/Gisson/simple-messenger/server"
	"github.com/julienschmidt/httprouter"
	_ "os"
)

type User struct {
	uid      int
	username string
	active   bool
}

type Message struct {
	owner   *User
	message string
	visible bool
	length  int
}

type MessageQueue struct {
	messages []*Message
}

type MessengerError struct {
	errorMessage string
}

func main() {
	server, _ := server.New(msgpkg.New(), httprouter.New())
	api.AddAllRoutes(server)
	server.Start()
}

func createRootUser() User {
	root := User{uid: 0, username: "root", active: true}
	return root
}

func addMessage(queue *MessageQueue, m Message) {
	queue.messages = append(queue.messages, &m)
	return
}

func getUserById(userList []User, uid int) (*User, error) {
	for _, user := range userList {
		if user.uid == uid {
			return &user, nil

		}
	}
	return nil, errors.New("No such uid")
}

func getUserMessages(queue *MessageQueue, uid int) []Message {
	userMessages := make([]Message, 0)
	for _, message := range queue.messages {
		if message.owner.uid == uid {
			userMessages = append(userMessages, *message)
		}
	}
	return userMessages
}

func getLongestMessage(queue *MessageQueue) []Message {
	longestMessages := make([]Message, 0)
	longestSize := 0
	for _, message := range queue.messages {
		if longestSize == message.length {
			longestMessages = append(longestMessages, *message)
		} else if message.length > longestSize {
			longestSize = message.length
			longestMessages = make([]Message, 1)
			longestMessages[0] = *message
		}
	}
	return longestMessages
}
