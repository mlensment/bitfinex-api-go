package rest

import (
	"fmt"
	"path"

	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2"
)

type TickerService struct {
	Synchronous
}

func (t *TickerService) Ticker(symbol string) (ticker *bitfinex.Ticker, err error) {
	if symbol == "" {
		err = fmt.Errorf("symbol cannot be empty")
		return
	}

	req := NewRequestWithMethod(path.Join("ticker", symbol), "GET")
	raw, err := t.Request(req)
	if err != nil {
		return
	}

	ticker, err = bitfinex.NewTickerFromRaw(symbol, raw)
	if err != nil {
		return
	}

	return
}
