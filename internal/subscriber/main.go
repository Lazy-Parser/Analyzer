package subscriber

import (
	"context"
	"flag"
	"log"

	"github.com/Lazy-Parser/Analyzer/internal/dispatcher"
	"github.com/nats-io/nats.go"
)

var (
	mexcSub = flag.String("mexcSub", "mexc.spread", "String subscription to NATS mexc SPOT / FUTURES")
)

func Start(conn *nats.Conn, d *dispatcher.Dispatcher) error {
	_, err := conn.Subscribe(*mexcSub, func(msg *nats.Msg) {
		go func() {
			err := d.Dispatch(context.Background(), msg)
			if err != nil {
				log.Printf("failed to dispatch message: %v", err)
			}
		}()
	})

	return err
}
