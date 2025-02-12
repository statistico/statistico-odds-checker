package bet365

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
)

var exchangeID = 2

type marketFactory struct {
	parser sportmonks.OddsParser
}

func (*marketFactory) Exchange() string {
	return "BET365"
}

func (m *marketFactory) CreateMarket(ctx context.Context, e *exchange.Event) (*exchange.Market, error) {
	runners, err := m.parser.ParseMarketRunners(ctx, int(e.ID), exchangeID, e.Market)

	if err != nil {
		return nil, err
	}

	if len(runners) == 0 {
		return nil, &exchange.NoEventMarketError{
			Exchange: "BET365",
			Market:   e.Market,
			EventID:  e.ID,
		}
	}

	return &exchange.Market{
		ID:       fmt.Sprintf("BET365-%d-%s", e.ID, e.Market),
		Name:     e.Market,
		EventID:  e.ID,
		Exchange: "BET365",
		Runners:  runners,
	}, nil
}

func NewMarketFactory(p sportmonks.OddsParser) exchange.MarketFactory {
	return &marketFactory{parser: p}
}
