-- +goose Up
-- +goose StatementBegin
CREATE TABLE market (
    event_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    side VARCHAR NOT NULL,
    exchange_odds JSONB NOT NULL,
    statistico_odds JSONB NOT NULL,
    timestamp INTEGER NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market;
-- +goose StatementEnd
