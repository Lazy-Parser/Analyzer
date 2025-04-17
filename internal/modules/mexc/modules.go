package mexc

type MexcModule struct{}

// "symbol":"SUI/USDT","spot_price":"2.1164","futures_price":"2.1152","timestamp":1744742029197
type MexcSpreadEvent struct {
	Symbol    string      `json:"symbol"`
	Futures   FuturesData `json:"futures"`
	Spot      SpotData    `json:"spot"`
	Timestamp int64       `json:"timestamp"`
}

type SpotData struct {
	Symbol     string  `json:"s"` // Kept as string
	Price      float64 `json:"p"`
	Change     float64 `json:"r"`
	TrueChange float64 `json:"tr"`
	High       float64 `json:"h"`
	Low        float64 `json:"l"`
	VolumeUSDT float64 `json:"v"`
	VolumeBase float64 `json:"q"`
	LastRT     float64 `json:"lastRT"`
	MT         float64 `json:"MT"`
	NV         string  `json:"NV"`
}

type FuturesData struct {
	Symbol       string  `json:"symbol"`
	LastPrice    float64 `json:"lastPrice"`
	RiseFallRate float64 `json:"riseFallRate"`
	FairPrice    float64 `json:"fairPrice"`
	IndexPrice   float64 `json:"indexPrice"`
	Volume24     float64 `json:"volume24"`
	Amount24     float64 `json:"amount24"`
	MaxBidPrice  float64 `json:"maxBidPrice"`
	MinAskPrice  float64 `json:"minAskPrice"`
	Lower24Price float64 `json:"lower24Price"`
	High24Price  float64 `json:"high24Price"`
	Timestamp    int64   `json:"timestamp"`
}
