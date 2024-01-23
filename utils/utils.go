package utils

import (
	"encoding/json"
	"time"
)

type Message struct {
	Name    string   `json:"name"`
	Message string   `json:"message"`
	Date    string   `json:"date"`
	Clients []string `json:"clients"`
}

func MarshalJsonMessage(name string, msg string, users []string) ([]byte, error) {
	date := time.Now().Format("2006-01-02 15:04")

	newServerMessage := Message{
		Name:    name,
		Message: msg,
		Date:    date,
		Clients: users,
	}
	jsonMessage, err := json.Marshal(newServerMessage)

	if err != nil {
		return nil, err
	}

	return jsonMessage, nil
}

func DecodeMessage(msg []byte) (Message, error) {
	var decodedMsg Message
	err := json.Unmarshal(msg, &decodedMsg)
	return decodedMsg, err
}
