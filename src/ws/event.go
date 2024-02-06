package ws

import (
	"encoding/json"
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
	UserName string `json:"userName"`
	Message  string `json:"message"`
}
