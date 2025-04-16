package dispatcher

import (
	"context"
	"encoding/json"
)


type EventType string

const (
	MEXC EventType = "mexc.spread"
)

type HandlerFunc func(ctx context.Context, data json.RawMessage) error

type Dispatcher struct {
	handlers 	map[EventType]HandlerFunc
}

type Event struct {
	Subject string          `json:"subject"`
	Data 	json.RawMessage `json:"data"`
}