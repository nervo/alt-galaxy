package message

const (
	OutMsg   MessageType = iota
	ErrorMsg MessageType = iota
)

type MessageType int

type Message struct {
	MessageType MessageType
	Body        string
}
