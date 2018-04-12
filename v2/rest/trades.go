package rest

import (
	"path"

	"github.com/bitfinexcom/bitfinex-api-go/v2"
)

// TradeService manages the Trade endpoint.
type TradeService struct {
	Synchronous
	Authenticator
}

// All returns all orders for the authenticated account.
func (s *TradeService) All(symbol string, params ...map[string]interface{}) (ts []*bitfinex.RestTrade, err error) {
	p := ReadParams(params...)
	r, err := s.NewAuthenticatedPostRequest(path.Join("auth", "r", "trades", symbol, "hist"), p)
	if err != nil {
		return
	}

	raw, err := s.Request(r)
	if err != nil {
		return
	}

	ts, err = bitfinex.NewRestTradeSliceFromRaw(raw)
	return
}
