package postgres

import (
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-checker/internal/market"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) Insert(m *market.Market) error {
	builder := r.queryBuilder()

	e, err := json.Marshal(m.ExchangeMarket)

	if err != nil {
		return err
	}

	s, err := json.Marshal(m.StatisticoOdds)

	if err != nil {
		return err
	}

	_, err = builder.
		Insert("market").
		Columns(
		"event_id",
			"name",
			"exchange",
			"side",
			"exchange_market",
			"statistico_odds",
			"timestamp",
		).
		Values(
			m.EventID,
			m.Name,
			m.Exchange,
			m.Side,
			string(e),
			string(s),
			m.Timestamp.Unix(),
		).
		Exec()

	return err
}

func (r *MarketRepository) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.connection)
}

func NewMarketRepository(connection *sql.DB) *MarketRepository {
	return &MarketRepository{connection: connection}
}
