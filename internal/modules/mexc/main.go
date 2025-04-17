package mexc

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
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
	minPercentDiff = flag.Float64("minPercentDiff", 1, "Minimum percent spread for spot / futures")
	minVolume      = flag.Float64("minVolume", 1000000, "Minimum volume for pair") // 2.5M
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

	// check data
	spread, ok := filterAlgorithm(evt)
	if !ok {
		return nil
	}

	fmt.Printf(
		"Монета: %s | SPOT & FUTURES\n Спред: %.2f%%\n Price Spot: %.5f$\n Price Futures: %.5f$ \n\n",
		evt.Symbol, spread, evt.Spot.Price, evt.Futures.LastPrice,
	)
	// exec.Command("afplay", "/System/Library/Sounds/Pop.aiff").Run()
	counter++
	publish("bot.mexc.spread", evt)

	return nil
}

func filterAlgorithm(evt MexcSpreadEvent) (float64, bool) {
	// check volumes
	if evt.Spot.VolumeUSDT < *minVolume || evt.Futures.Amount24 < *minVolume {
		return 0, false
	}

	// check spread
	spread := findSpread(evt.Spot.Price, evt.Futures.LastPrice)
	if spread < *minPercentDiff {
		return 0, false
	}

	// different directions
	// spotChange := evt.Spot.Change
	// futuresChange := evt.Futures.RiseFallRate * 100
	// if (spotChange > 0 && futuresChange < 0) || (spotChange < 0 && futuresChange > 0) {
	// 	return 0, false
	// }

	return spread, true
}

func findSpread(a float64, b float64) float64 {
	return math.Abs(((a - b) / a) * 100)
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
