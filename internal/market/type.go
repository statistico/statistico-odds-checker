package market

import (
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc"
	"time"
)

type Market struct {
	EventID        uint64
	Name           string
	Side           string
	ExchangeName   string
	ExchangeMarket exchange.Market
	ImpliedOdds    []*grpc.Odds
	Timestamp      time.Time
}
