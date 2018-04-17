package main

import (
	"fmt"
	"log"

	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

func main() {
	c := rest.NewClient()
	ticker, err := c.Ticker.Ticker("tETHUSD")
	if err != nil {
		log.Fatalf("getting ticker: %s", err)
	}
	fmt.Printf("Symbol: %v Ask: %v Bid: %v", ticker.Symbol, ticker.Ask, ticker.Bid)
}
