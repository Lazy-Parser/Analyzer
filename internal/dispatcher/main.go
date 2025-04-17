package dispatcher

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

func New() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[EventType]HandlerFunc),
	}
}

func (d *Dispatcher) Register(tp EventType, handler HandlerFunc) {
	d.handlers[tp] = handler
}

func (d *Dispatcher) Dispatch(ctx context.Context, msg *nats.Msg) error {
	hanler, ok := d.handlers[EventType(msg.Subject)]
	if !ok {
		return fmt.Errorf("no handler for event type: %s", msg.Subject)
	}

	return hanler(ctx, msg.Data)
}
