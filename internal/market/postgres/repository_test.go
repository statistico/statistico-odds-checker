package postgres_test

import (
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/market"
	"github.com/statistico/statistico-odds-checker/internal/market/postgres"
	"github.com/statistico/statistico-odds-checker/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketRepository_Insert(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "market")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Market *market.Market
			Count int8
		}{
			{newMarket(182981, time.Now()), 1},
			{newMarket(182981, time.Now()), 2},
			{newMarket(182981, time.Now()), 3},
		}

		for _, tc := range tradeCounts {
			insertTrade(t, repo, tc.Market)

			var count int8

			row := conn.QueryRow("select count(*) from market")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func newMarket(eventID uint64, t time.Time) *market.Market {
	return &market.Market{
		EventID:        eventID,
		Name:           "OVER_UNDER_25",
		Side:           "BACK",
		Exchange:   "betfair",
		ExchangeMarket: exchange.Market{
			ID:      "1.2981871",
			Runners: []exchange.Runner{
				{
					ID: 48291,
					Name: "Over 2.5 Goals",
					Prices: []exchange.PriceSize{
						{
							Price: 1.95,
							Size: 159001,
						},
						{
							Price: 2.00,
							Size: 50.56,
						},
					},
				},
			},
		},
		Timestamp:      t,
	}
}

func insertTrade(t *testing.T, repo *postgres.MarketRepository, m *market.Market) {
	if err := repo.Insert(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}
