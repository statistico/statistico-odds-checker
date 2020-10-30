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
			{newMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 1},
			{newMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 2},
			{newMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 3},
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

func TestMarketRepository_Get(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "market")
	repo := postgres.NewMarketRepository(conn)

	t.Run("returns a slice of filtered market struct", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		timestamp := time.Date(2020, 03, 12, 0, 0, 0, 0, &time.Location{})

		marketOne := newMarket(1234, "OVER_UNDER_25", "BACK", timestamp)
		marketTwo := newMarket(5678, "OVER_UNDER_15", "BACK", timestamp)
		marketThree := newMarket(9999, "OVER_UNDER_25", "LAY", timestamp)

		insertTrade(t, repo, marketOne)
		insertTrade(t, repo, marketTwo)
		insertTrade(t, repo, marketThree)

		cases := []struct{
			Query *market.RepositoryQuery
			Markets []*market.Market
		} {
			{
				Query: &market.RepositoryQuery{
					EventID:    &marketOne.EventID,
				},
				Markets: []*market.Market{marketOne},
			},
			{
				Query: &market.RepositoryQuery{
					MarketName:    &marketOne.Name,
				},
				Markets: []*market.Market{marketOne, marketThree},
			},
			{
				Query: &market.RepositoryQuery{
					Side:    &marketThree.Side,
				},
				Markets: []*market.Market{marketThree},
			},
			{
				Query: &market.RepositoryQuery{
					EventID: &marketOne.EventID,
					MarketName: &marketTwo.Name,
					Side: &marketThree.Side,
				},
			},
		}

		for _, c := range cases {
			markets, err := repo.Get(c.Query)

			if err != nil {
				t.Errorf("Error fetching markets from the database: %s", err.Error())
			}

			assert.Equal(t, c.Markets, markets)
		}
	})
}

func newMarket(eventID uint64, name, side string, t time.Time) *market.Market {
	return &market.Market{
		EventID:        eventID,
		Name:           name,
		Side:           side,
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
