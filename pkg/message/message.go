package message

type MessageCategory string

// Message categories
const SIGNAL MessageCategory = "SIGNAL"
const ID MessageCategory = "ID"
const CHAT MessageCategory = "CHAT"

// Signal connection messages
const NEW_CONNECTION_SIGNAL string = "NEW_CONNECTION_SIGNAL"
const CONNECTED_SIGNAL string = "CONNECTED_SIGNAL"
const DISCONNECTED_SIGNAL string = "DISCONNECTED_SIGNAL"

type Message struct {
	To          string          `json:"to"`
	Category    string          `json:"category,omitempty"`
	MessageType MessageCategory `json:"messageType"`
	Content     string          `json:"content"`
}
