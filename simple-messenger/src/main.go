package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	root := createRootUser()
	userList := make([]User, 1)
	userList[0] = root
	queue := MessageQueue{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
		if text == "A\n" {
			var uid int
			var messageText string

			_, err := fmt.Scanf("%d %s", &uid, &messageText)
			if err != nil {
				fmt.Print(err)
				os.Exit(-1)
			}
			user, err := getUserById(userList, uid)
			if err != nil {
				fmt.Print(err)
			}
			message := Message{owner: user, message: messageText, visible: true, length: len(messageText)}
			addMessage(&queue, message)
		} else if text == "L\n" {
			fmt.Printf("Total messages: %d\n", len(queue.messages))
			for _, message := range queue.messages {
				fmt.Printf("%d:%s\n", message.owner.uid, message.message)
			}
		} else if text == "U\n" {
			var uid int
			fmt.Scanf("%d", &uid)
			fmt.Printf("Messages from user %d\n", uid)
			for _, message := range getUserMessages(&queue, uid) {
				fmt.Println(message.message)
			}
		} else if text == "O\n" {
			for _, message := range getLongestMessage(&queue) {
				fmt.Printf("Longest sentence:%s:%s\n", message.owner.username, message.message)
			}
		} else if text == "X\n" {
			break
		}
	}
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
