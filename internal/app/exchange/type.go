package exchange

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Query struct {
	Event  string
	Date   time.Time
	Market string
	Sport  string
}

type Market struct {
	ID           string    `json:"id"`
	ExchangeName string    `json:"exchange"`
	Side         string    `json:"side"`
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
	ID     uint64      `json:"id"`
	Name   string      `json:"name"`
	Sort   int         `json:"sort"`
	Prices []PriceSize `json:"prices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
	Side  string  `json:"side"`
}
