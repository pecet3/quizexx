package ws

import (
	"encoding/json"
	"time"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *client) error

const (
	EventSendMessage = "send_message"
	EventSendAnswer  = "send_answer"
)

type SendMessageEvent struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type SendAnswerEvent struct {
	UserName string    `json:"user_name"`
	AnswerId uint8     `json:"answer_id"`
	Time     time.Time `json:"time"`
}
