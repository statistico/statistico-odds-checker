package sport

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"time"
)

type EventMarketRequester interface {
	FindEventMarkets(ctx context.Context, from, to time.Time) <-chan *EventMarket
}

type EventMarket struct {
	EventID        uint64          `json:"eventId"`
	Sport          string          `json:"sport"`
	EventDate      int64           `json:"eventDate"`
	MarketName     string          `json:"name"`
	Side           string          `json:"side"`
	Exchange       string          `json:"exchange"`
	ExchangeMarket exchange.Market `json:"exchangeMarket"`
	StatisticoOdds []*proto.Odds   `json:"statisticoOdds"`
	Timestamp      int64           `json:"timestamp"`
}
