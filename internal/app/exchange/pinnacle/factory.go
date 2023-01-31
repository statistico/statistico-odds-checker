package pinnacle

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
)

var exchangeID = 70

type marketFactory struct {
	parser sportmonks.OddsParser
}

func (m *marketFactory) CreateMarket(ctx context.Context, e *exchange.Event) (*exchange.Market, error) {
	odds, err := m.parser.ParseMarketOdds(ctx, int(e.ID), exchangeID, e.Market)

	if err != nil {
		return nil, err
	}

	runners, err := exchange.ConvertOddsToRunners(odds)

	if err != nil {
		return nil, fmt.Errorf("error converting odds to runners in Pinnacle exchange: %s", err.Error())
	}

	return &exchange.Market{
		ID:       "NONE",
		Name:     e.Market,
		Exchange: "PINNACLE",
		Runners:  runners,
	}, nil
}

func NewMarketFactory(p sportmonks.OddsParser) exchange.MarketFactory {
	return &marketFactory{parser: p}
}
