package message

import (
	"errors"
	"fmt"
	"sort"
)

type MessageManager struct {
	messages  []Message
	currentid int
}

func New() *MessageManager {
	manager := &MessageManager{currentid: 0}
	manager.messages = make([]Message, 0)
	return manager
}

func (manager *MessageManager) NewMessage(message string, owner int, length int) {
	newmessage := Message{uid: manager.currentid, Message: message, visible: true, Owner: owner, length: len(message)}
	manager.messages = append(manager.messages, newmessage)
	manager.currentid++

}

func (manager *MessageManager) AddMessage(message Message) {
	fmt.Printf("Message uid: %d\n", message.uid)
	fmt.Printf("Currentuid: %d\n", manager.currentid)
	message.uid = manager.currentid
	message.visible = true
	message.length = len(message.Message)
	manager.messages = append(manager.messages, message)
	manager.currentid++
}

func (manager *MessageManager) DeleteMessage(messageid int) error {
	for i, message := range manager.messages {
		if message.uid == messageid {
			manager.messages = append(manager.messages[:i], manager.messages[i+1:]...)
			return nil
		}
	}
	return errors.New("No such uid")

}

func (manager *MessageManager) ListMessages() []Message {
	return manager.messages
}

func (manager *MessageManager) ListUserMessages(uid int) []Message {
	userMessages := make([]Message, 0)
	for _, message := range manager.messages {
		if message.Owner == uid {
			userMessages = append(userMessages, message)
		}
	}
	return userMessages
}

func (manager *MessageManager) GetLongestSentence() []Message {
	longestMessages := make([]Message, 0)
	longestSize := 0
	for _, message := range manager.messages {
		if message.length == longestSize {
			longestMessages = append(longestMessages, message)
		} else if message.length > longestSize {
			longestMessages := make([]Message, 1)
			longestMessages[0] = message
			longestSize++
		}
	}
	return longestMessages
}

func (manager *MessageManager) GetSortedMessages() []Message {
	sorted := make([]Message, len(manager.messages))
	copy(sorted, manager.messages)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].length > sorted[j].length
	})
	return sorted
}

func (manager *MessageManager) GetNumberOfMessages() int {
	return len(manager.messages)
}

/*func (manager *MessageManager) GetMostActiveUser() ([int]map[int]) {
}TODO*/
