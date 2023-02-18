package bootstrap

import (
	bfc "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/bet365"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betcris"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betfair"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/pinnacle"
	"net/http"
)

func (c Container) BetfairMarketFactory() exchange.MarketFactory {
	config := c.Config.BetFair

	creds := bfc.InteractiveCredentials{
		Username: config.Username,
		Password: config.Password,
		Key:      config.Key,
	}

	urls := bfc.BaseURLs{
		Accounts: bfc.AccountsURL,
		Betting:  bfc.BettingURL,
		Login:    bfc.LoginURL,
	}

	client := bfc.Client{
		HTTPClient:  &http.Client{},
		Credentials: creds,
		BaseURLs:    urls,
	}
	return betfair.NewMarketFactory(client)
}

func (c Container) Bet365MarketFactory() exchange.MarketFactory {
	return bet365.NewMarketFactory(c.SportmonksOddsParser())
}

func (c Container) BetCrisMarketFactory() exchange.MarketFactory {
	return betcris.NewMarketFactory(c.SportmonksOddsParser())
}

func (c Container) PinnacleMarketFactory() exchange.MarketFactory {
	return pinnacle.NewMarketFactory(c.SportmonksOddsParser())
}

func (c Container) MarketBuilder() exchange.MarketBuilder {
	factories := []exchange.MarketFactory{
		//c.Bet365MarketFactory(),
		//c.BetCrisMarketFactory(),
		c.PinnacleMarketFactory(),
	}

	return exchange.NewMarketBuilder(factories, c.Logger)
}
