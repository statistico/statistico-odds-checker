package market

import (
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
)

type Market struct {
	EventID        uint64
	Name           string
	ExchangeMarket exchange.Market
	ImpliedOdds    []*proto.Odds
}
