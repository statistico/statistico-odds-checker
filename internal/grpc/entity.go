package grpc

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Odds struct {
	Price                float32  `json:"price"`
	Selection            string   `json:"selection"`
}

func (o Odds) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Odds) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}
