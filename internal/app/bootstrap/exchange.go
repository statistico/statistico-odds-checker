package bootstrap

import (
	bfc "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/bet365"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betfair"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/statistico"
	"net/http"
)

func (c Container) BetfairMarketFactory() exchange.MarketFactory {
	config := c.Config.BetFair

	creds := bfc.InteractiveCredentials{
		Username: config.Username,
		Password: config.Password,
		Key:      config.Key,
	}

	store := c.Cache()

	client := bfc.NewClient(&http.Client{}, creds, store)

	return betfair.NewMarketFactory(client)
}

func (c Container) Bet365MarketFactory() exchange.MarketFactory {
	return bet365.NewMarketFactory(c.SportmonksOddsParser())
}

func (c Container) StatisticoMarketFactory() exchange.MarketFactory {
	return statistico.NewMarketFactory(c.OddsCompilerClient())
}

func (c Container) MarketFactoryResolver() exchange.MarketFactoryResolver {
	factories := []exchange.MarketFactory{
		c.Bet365MarketFactory(),
		//c.BetfairMarketFactory(),
		//c.StatisticoMarketFactory(),
	}

	return exchange.NewMarketFactoryResolver(factories)
}
