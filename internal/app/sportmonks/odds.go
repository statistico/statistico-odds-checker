package sportmonks

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

type OddsParser interface {
	// ParseMarketOdds parses and returns a slice of sportmonks.Odds struct associated to a fixture, exchange
	// and market.
	ParseMarketOdds(ctx context.Context, fixtureID, exchangeID int, market string) ([]sportmonks.Odds, error)
}

type oddsParser struct {
	client *sportmonks.HTTPClient
}

func (m *oddsParser) ParseMarketOdds(ctx context.Context, fixtureID, exchangeID int, market string) ([]sportmonks.Odds, error) {
	marketID, ok := marketIDs[market]

	if !ok {
		return nil, fmt.Errorf("error handling market for exchange '%d': market '%s' is not supported", exchangeID, market)
	}

	markets, _, err := m.client.OddsByFixtureAndMarket(ctx, fixtureID, marketID, []string{"bookmaker", "market"})

	if err != nil {
		return nil, fmt.Errorf("error fetching markets for exchange '%d': %s", exchangeID, err.Error())
	}

	if len(markets) == 0 {
		return []sportmonks.Odds{}, nil
	}

	odds := parseExchangeMarketOdds(exchangeID, markets)

	if odds == nil || len(odds) == 0 {
		return []sportmonks.Odds{}, nil
	}

	return parseMarketRunners(market, exchangeID, odds)
}

func NewOddsParser(c *sportmonks.HTTPClient) OddsParser {
	return &oddsParser{client: c}
}
