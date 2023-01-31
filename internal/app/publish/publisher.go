package publish

import (
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
)

type Publisher interface {
	PublishMarket(m *stream.EventMarket) error
}
