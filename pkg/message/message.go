package message

type MessageCategory string

// Message categories
const SIGNAL MessageCategory = "SIGNAL"
const ID MessageCategory = "ID"
const CHAT MessageCategory = "CHAT"
const OFFER MessageCategory = "OFFER"

// Signal connection messages
const CONNECT_REQ string = "CONNECT_REQ"
const DISCONNECT_REQ string = "DISCONNECT_REQ"
const SKIP_REQ string = "SKIP_REQ"
const CONNECT_FAIL string = "CONNECTION_FAIL"
const CONNECT_SIGNAL string = "CONNECT_SIGNAL"
const DISCONNECT_SIGNAL string = "DISCONNECT_SIGNAL"
const ICE_SIGNAL string = "ICE_SIGNAL"

// Offer connection category
const OFFER_REQ string = "OFFER_REQ"
const OFFER_ACC string = "OFFER_ACC"

type Message struct {
	From        string          `json:"from,omitempty"`
	To          string          `json:"to"`
	Category    string          `json:"category,omitempty"`
	MessageType MessageCategory `json:"messageType"`
	Content     interface{}     `json:"content"`
}
