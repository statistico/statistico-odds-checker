package sportmonks

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
)

type OddsParser interface {
	// ParseMarketRunners parses and returns a slice of exchange.Runner struct associated to a fixture, exchange
	// and market.
	ParseMarketRunners(ctx context.Context, fixtureID, exchangeID int, market string) ([]*exchange.Runner, error)
}

type oddsParser struct {
	client *sportmonks.HTTPClient
}

func (m *oddsParser) ParseMarketRunners(ctx context.Context, fixtureID, exchangeID int, market string) ([]*exchange.Runner, error) {
	marketID, ok := marketIDs[market]

	if !ok {
		return nil, fmt.Errorf("error handling market for exchange '%d': market '%s' is not supported", exchangeID, market)
	}

	markets, _, err := m.client.OddsByFixtureAndMarket(ctx, fixtureID, marketID, []string{"bookmaker", "market"})

	if err != nil {
		return nil, fmt.Errorf("error fetching markets for exchange '%d': %s", exchangeID, err.Error())
	}

	if len(markets) == 0 {
		return []*exchange.Runner{}, nil
	}

	odds := parseExchangeMarketOdds(exchangeID, markets)

	if odds == nil || len(odds) == 0 {
		return []*exchange.Runner{}, nil
	}

	return convertOddsToRunners(odds, market)
}

func NewOddsParser(c *sportmonks.HTTPClient) OddsParser {
	return &oddsParser{client: c}
}
