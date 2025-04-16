package main

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/Analyzer/internal/dispatcher"
	"github.com/Lazy-Parser/Analyzer/internal/modules/mexc"
	"github.com/Lazy-Parser/Analyzer/internal/subscriber"
	"github.com/Lazy-Parser/Analyzer/internal/utils"
	"github.com/nats-io/nats.go"
)

func main() {
	natsUrl, err := utils.GetNatsUrl()
	if err != nil {
		fmt.Errorf("Get NATS URL: %w", err)
	}

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