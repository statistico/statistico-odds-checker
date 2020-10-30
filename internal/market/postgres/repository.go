package postgres

import (
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-checker/internal/market"
	"time"
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

func (r *MarketRepository) Get(q *market.RepositoryQuery) ([]*market.Market, error) {
	builder := r.queryBuilder()

	query := builder.Select("market.*").From("market")

	rows, err := buildQuery(query, q).Query()

	if err != nil {
		return []*market.Market{}, err
	}

	return rowsToMarketSlice(rows)
}

func buildQuery(b sq.SelectBuilder, q *market.RepositoryQuery) sq.SelectBuilder {
	if q.EventID != nil {
		b = b.Where(sq.Eq{"event_id": *q.EventID})
	}

	if q.MarketName != nil {
		b = b.Where(sq.Eq{"name": *q.MarketName})
	}

	if q.Side != nil {
		b = b.Where(sq.Eq{"side": *q.Side})
	}

	if q.SortBy != nil && *q.SortBy == "timestamp_asc" {
		b = b.OrderBy("timestamp ASC")
	}

	if q.SortBy != nil && *q.SortBy == "timestamp_desc" {
		b = b.OrderBy("timestamp DESC")
	}

	if q.SortBy == nil {
		b = b.OrderBy("timestamp DESC")
	}

	return b
}

func rowsToMarketSlice(rows *sql.Rows) ([]*market.Market, error) {
	defer rows.Close()

	var markets []*market.Market

	for rows.Next() {
		var m market.Market
		var exchange string
		var odds string
		var timestamp int64

		err := rows.Scan(
			&m.EventID,
			&m.Name,
			&m.Exchange,
			&m.Side,
			&exchange,
			&odds,
			&timestamp,
		)

		if err != nil {
			return markets, err
		}

		m.Timestamp = time.Unix(timestamp, 0)

		err = json.Unmarshal([]byte(exchange), &m.ExchangeMarket)

		if err != nil {
			return markets, err
		}

		err = json.Unmarshal([]byte(odds), &m.StatisticoOdds)

		if err != nil {
			return markets, err
		}

		markets = append(markets, &m)
	}

	return markets, nil
}

func (r *MarketRepository) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.connection)
}

func NewMarketRepository(connection *sql.DB) *MarketRepository {
	return &MarketRepository{connection: connection}
}
