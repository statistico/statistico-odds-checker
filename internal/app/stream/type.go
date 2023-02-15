package stream

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"time"
)

type EventMarketStreamer interface {
	Stream(ctx context.Context, from, to time.Time) <-chan *EventMarket
}

type EventMarket struct {
	ID            string             `json:"id"`
	EventID       uint64             `json:"eventId"`
	CompetitionID uint64             `json:"competitionId"`
	SeasonID      uint64             `json:"seasonId"`
	EventDate     int64              `json:"eventDate"`
	MarketName    string             `json:"name"`
	Exchange      string             `json:"exchange"`
	Runners       []*exchange.Runner `json:"runners"`
	Timestamp     int64              `json:"timestamp"`
}
