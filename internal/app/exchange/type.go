package exchange

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Event struct {
	ID     uint64
	Name   string
	Date   time.Time
	Market string
}

type Market struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ExchangeName string    `json:"exchange"`
	Runners      []*Runner `json:"runners"`
}

func (m Market) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Market) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

type Runner struct {
	ID         uint64      `json:"id"`
	Name       string      `json:"name"`
	BackPrices []PriceSize `json:"backPrices"`
	LayPrices  []PriceSize `json:"layPrices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}

type MarketFactory interface {
	CreateMarket(ctx context.Context, q *Event) (*Market, error)
}
