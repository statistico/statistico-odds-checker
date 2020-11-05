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
	ID              string             `json:"id"`
	EventID         uint64             `json:"eventId"`
	Sport           string             `json:"sport"`
	EventDate       int64              `json:"eventDate"`
	MarketName      string             `json:"name"`
	Side            string             `json:"side"`
	Exchange        string             `json:"exchange"`
	ExchangeRunners []*exchange.Runner `json:"exchangeRunners"`
	StatisticoOdds  []*proto.Odds      `json:"statisticoOdds"`
	Timestamp       int64              `json:"timestamp"`
}
