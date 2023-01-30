package bootstrap

import (
	bfc "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betfair"
	"net/http"
)

func (c Container) MarketFactory() exchange.MarketFactory {
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
