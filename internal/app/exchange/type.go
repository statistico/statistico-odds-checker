package exchange

import (
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
	ID       string    `json:"id"`
	Exchange string    `json:"exchange"`
	Name     string    `json:"name"`
	EventID  uint64    `json:"event_id"`
	Runners  []*Runner `json:"runners"`
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
	Name       *string     `json:"name"`
	Label      string      `json:"label"`
	Value      *float64    `json:"value"`
	BackPrices []PriceSize `json:"backPrices"`
	LayPrices  []PriceSize `json:"layPrices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
