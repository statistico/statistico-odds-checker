package postgres_test

import (
	"github.com/google/uuid"
	"github.com/statistico/statistico-odds-checker/internal/market/postgres"
	"github.com/statistico/statistico-odds-checker/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTradeRepository_InsertTrade(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "trade")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Trade *trade.Trade
			Count int8
		}{
			{newTrade("123"), 1},
			{newTrade("456"), 2},
			{newTrade("789"), 3},
		}

		for _, tc := range tradeCounts {
			insertTrade(t, repo, tc.Trade)

			var count int8

			row := conn.QueryRow("select count(*) from trade")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func TestTradeRepository_InsertTradeStatus(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "trade_status")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increase table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Trade *trade.Status
			Count int8
		}{
			{newTradeStatus("pending"), 1},
			{newTradeStatus("won"), 2},
			{newTradeStatus("lost"), 3},
		}

		for _, tc := range tradeCounts {
			insertTradeStatus(t, repo, tc.Trade)

			var count int8

			row := conn.QueryRow("select count(*) from trade_status")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func newTrade(transactionID string) *trade.Trade {
	return &trade.Trade{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		TransactionID:  transactionID,
		Name:           "Over 2.5 Goals",
		Side:           "BACK",
		Exchange:       "betfair",
		Market:         "OVER_UNDER_25",
		EventID:        "12345",
		ExchangeOdds:   2.10,
		StatisticoOdds: 1.45,
		Stake:          50,
		Strategy:       "half_kelly",
		Timestamp:      time.Now(),
	}
}

func insertTrade(t *testing.T, repo *postgres.TradeRepository, trade *trade.Trade) {
	if err := repo.InsertTrade(trade); err != nil {
		t.Errorf("Error when inserting trade into the database: %s", err.Error())
	}
}

func newTradeStatus(status string) *trade.Status {
	return &trade.Status{
		TradeID:   uuid.UUID{},
		Status:    status,
		Timestamp: time.Now(),
	}
}

func insertTradeStatus(t *testing.T, repo *postgres.TradeRepository, status *trade.Status) {
	if err := repo.InsertTradeStatus(status); err != nil {
		t.Errorf("Error when inserting trade status into the database: %s", err.Error())
	}
}
