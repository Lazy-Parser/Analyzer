package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lazy-Parser/Analyzer/internal/dispatcher"
	"github.com/Lazy-Parser/Analyzer/internal/modules/mexc"
	"github.com/Lazy-Parser/Analyzer/internal/subscriber"
	"github.com/nats-io/nats.go"
)

func main() {
	natsUrl := os.Getenv("NATS_URL")

	conn, err := nats.Connect(natsUrl)
	if err != nil {
		fmt.Errorf("Get NATS URL: %w", err)
		return
	}
	fmt.Println("Connected to NATS! (Subscribe) ")
	defer conn.Drain()

	d := dispatcher.New()

	mexc.New(conn).Register(d)

	if err := subscriber.Start(conn, d); err != nil {
		log.Fatal("subscriber error:", err)
	}

	select {}
}