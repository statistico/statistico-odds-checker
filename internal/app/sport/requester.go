package sport

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"time"
)

type EventMarketRequester interface {
	FindEventMarkets(ctx context.Context, from, to time.Time) <-chan *EventMarket
}

type EventMarket struct {
	ID              string             `json:"id"`
	EventID         uint64             `json:"eventId"`
	CompetitionID   uint64             `json:"competitionId"`
	SeasonID        uint64             `json:"seasonId"`
	Sport           string             `json:"sport"`
	EventDate       string             `json:"date"`
	MarketName      string             `json:"name"`
	Side            string             `json:"side"`
	Exchange        string             `json:"exchange"`
	Runners         []*exchange.Runner `json:"runners"`
	Timestamp       int64              `json:"timestamp"`
}
