package network

type EventType string

type Event interface {
	TypeOf() EventType
}
