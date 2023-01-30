package sportmonks

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

type OddsParser interface {
	ParseMarketOdds(ctx context.Context, fixtureID, exchangeID int, market string) ([]sportmonks.Odds, error)
}

type oddsParser struct {
	client *sportmonks.HTTPClient
}

func (m *oddsParser) ParseMarketOdds(ctx context.Context, fixtureID, exchangeID int, market string) ([]sportmonks.Odds, error) {
	marketId, err := parseMarketId(market)

	if err != nil {
		return nil, fmt.Errorf("error handling market for exchange '%d': %s", exchangeID, err.Error())
	}

	markets, _, err := m.client.OddsByFixtureAndMarket(ctx, fixtureID, marketId)

	if err != nil {
		return nil, fmt.Errorf("error fetching markets for exchange '%d': %s", exchangeID, err.Error())
	}

	if len(markets) == 0 {
		return nil, nil
	}

	odds := parseExchangeOdds(exchangeID, markets)

	if odds == nil || len(odds) == 0 {
		return nil, nil
	}

	return parseMarketRunners(market, exchangeID, odds)
}

func NewOddsParser(c *sportmonks.HTTPClient) OddsParser {
	return &oddsParser{client: c}
}
