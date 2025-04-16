package mexc

// "symbol":"SUI/USDT","spot_price":"2.1164","futures_price":"2.1152","timestamp":1744742029197
type MexcSpreadEvent struct {
	Symbol       string  `json:"symbol"`
	PriceSpot    string  `json:"spot_price"`
	PriceFutures string  `json:"futures_price"`
	Amount24     float64 `json:"amount24"`
	Timestamp    int64   `json:"timestamp"`
}

type MexcModule struct{}
