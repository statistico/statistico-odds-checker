package log

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	"github.com/statistico/statistico-odds-checker/internal/app/sport"
)

type Publisher struct {
	logger *logrus.Logger
}

func (p *Publisher) PublishMarket(m *sport.EventMarket) error {
	p.logger.Infof("Pretending to publish market for event %d and market %s", m.EventID, m.MarketName)

	return nil
}

func NewPublisher(l *logrus.Logger) publish.Publisher {
	return &Publisher{logger: l}
}
