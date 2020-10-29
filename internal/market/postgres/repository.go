package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-checker/internal/market"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) Insert(m *market.Market) error {
	builder := r.queryBuilder()

	_, err := builder.
		Insert("trade").
		Columns(
		"event_id",
			"name",
			"exchange",
			"side",
			"exchange_odds",
			"statistico_odds",
			"timestamp",
		).
		Values(
			m.EventID,
			m.Name,
			m.ExchangeName,
			m.Side,
			m.ExchangeMarket,
			m.ImpliedOdds,
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
