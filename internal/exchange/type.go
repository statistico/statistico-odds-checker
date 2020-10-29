package exchange

import "time"

type Query struct {
	Event  string
	Date   time.Time
	Market string
	Sport  string
}

type Market struct {
	ID      string
	Runners []Runner
}

type Runner struct {
	ID     uint64
	Name   string
	Prices []PriceSize
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}

type Ticket struct {
	MarketID        string
	MarketName      string
	SelectionID     uint64
	SelectionName   string
	Side            string
	StatisticoPrice float32
	ExchangePrice   float32
	Size            float32
	PersistenceType string
	OrderType       string
	Strategy        string
}

type Transaction struct {
	Exchange    string
	BetID       string
	MarketID    string
	SelectionID uint64
	Name        string
	Side        string
	Price       float32
	Size        float32
	OrderType   string
	Status      string
}
