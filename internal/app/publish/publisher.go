package publish

import "github.com/statistico/statistico-odds-checker/internal/app/sport"

type Publisher interface {
	PublishMarket(m *sport.EventMarket) error
}
