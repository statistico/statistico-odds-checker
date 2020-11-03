package market

import (
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
)

type Market struct {
	EventID        uint64          `json:"event_id"`
	Name           string          `json:"name"`
	Side           string          `json:"side"`
	Exchange       string          `json:"exchange"`
	ExchangeMarket exchange.Market `json:"exchange_market"`
	StatisticoOdds []*proto.Odds   `json:"statistico_odds"`
}
