package mexc

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"

	"github.com/Lazy-Parser/Analyzer/internal/dispatcher"
	"github.com/nats-io/nats.go"
)

// type Message struct {
// 	Symbol       string `json:"symbol"`
// 	SpotPrice    string `json:"spot_price"`
// 	FuturesPrice string `json:"futures_price"`
// 	Timestamp    int64  `json:"timestamp"`
// }

var (
	counter        = 1
	conn           *nats.Conn
	minPercentDiff = flag.Float64("minPercentDiff", 3, "Minimum percent spread for spot / futures")
)

func New(c *nats.Conn) *MexcModule {
	conn = c
	return &MexcModule{}
}

func (m *MexcModule) Register(d *dispatcher.Dispatcher) {
	d.Register("mexc.spread", m.handleSpread)
}

func (m *MexcModule) handleSpread(ctx context.Context, data json.RawMessage) error {

	var evt MexcSpreadEvent
	if err := json.Unmarshal(data, &evt); err != nil {
		return fmt.Errorf("parse data from mexc.spread failed: %w", err)
	}

	if evt.Amount24 < 1000000 {
		// fmt.Printf("Too litle: %f \n", evt.Amount24)
		return nil
	}

	// main logic
	f := toFloat(evt.PriceFutures)
	s := toFloat(evt.PriceSpot)
	percentDiff := math.Abs(((f - s) / f) * 100)

	if math.Abs(percentDiff) >= *minPercentDiff {
		log.Printf("%d) DIFF %f! %s.\n spot: %s  | futures: %s | amount24: %f \n\n", counter, percentDiff, evt.Symbol, evt.PriceSpot, evt.PriceFutures, evt.Amount24)
		exec.Command("afplay", "/System/Library/Sounds/Pop.aiff").Run()
		counter++
		publish("bot.mexc.spread", evt)
	}

	return nil
}

func toFloat(data string) float64 {
	res, _ := strconv.ParseFloat(data, 64)
	return res
}

func publish(subject string, data MexcSpreadEvent) {
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("marshal nats payload: %w", err)
	}

	err = conn.Publish(subject, payload)
	if err != nil {
		fmt.Errorf("Publish message: %w", err)
	}

	log.Println("New message published!")
}
