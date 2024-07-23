package message

type MessageCategory string

// Message categories
const SIGNAL MessageCategory = "SIGNAL"
const ID MessageCategory = "ID"
const TEXT_MESSAGE MessageCategory = "TEXT"

type Chat struct {
	To          string          `json:"to"`
	MessageType MessageCategory `json:"messageType"`
	Content     string          `json:"content"`
}
