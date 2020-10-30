package market

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc"
	"time"
)

type Market struct {
	EventID        uint64          `json:"event_id"`
	Name           string          `json:"name"`
	Side           string          `json:"side"`
	Exchange       string          `json:"exchange"`
	ExchangeMarket exchange.Market `json:"exchange_market"`
	StatisticoOdds []*grpc.Odds    `json:"statistico_odds"`
	Timestamp      time.Time       `json:"timestamp"`
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
