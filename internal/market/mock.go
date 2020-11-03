package market

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockMarketBuilder struct {
	mock.Mock
}

func (m *MockMarketBuilder) Build(ctx context.Context, q *BuilderQuery) <-chan *Market {
	args := m.Called(ctx, q)
	return args.Get(0).(<-chan *Market)
}


