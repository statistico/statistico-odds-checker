package market

import (
	"github.com/statistico/statistico-odds-checker/internal/exchange"
)

type Market struct {
	ID              string             `json:"id"`
	Exchange        string             `json:"exchange"`
	EventID         uint64             `json:"event_id"`
	Name            string             `json:"name"`
	Side            string             `json:"side"`
	ExchangeRunners []*exchange.Runner `json:"exchange_runners"`
}
