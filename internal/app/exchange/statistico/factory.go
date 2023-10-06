package statistico

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-compiler-go-grpc-client"
	"strings"
)

type marketFactory struct {
	compiler statisticooddscompiler.OddCompilerClient
}

func (m *marketFactory) Exchange() string {
	return "STATISTICO"
}

func (m *marketFactory) CreateMarket(ctx context.Context, e *exchange.Event) (*exchange.Market, error) {
	if supported := marketIsSupported(e.Market); !supported {
		return nil, nil
	}

	market, err := m.compiler.GetEventMarket(ctx, e.ID, e.Market)

	if err != nil {
		return nil, err
	}

	em := exchange.Market{
		ID:       fmt.Sprintf("STA-%d-%s", e.ID, e.Market),
		Exchange: "STATISTICO",
		Name:     e.Market,
		EventID:  e.ID,
	}

	for _, o := range market.Odds {
		r := exchange.Runner{
			Name: strings.ToUpper(o.Selection),
			BackPrices: []exchange.PriceSize{
				{
					Price: o.Price,
					Size:  0,
				},
			},
		}

		em.Runners = append(em.Runners, &r)
	}

	return &em, nil
}

func NewMarketFactory(c statisticooddscompiler.OddCompilerClient) exchange.MarketFactory {
	return &marketFactory{compiler: c}
}
