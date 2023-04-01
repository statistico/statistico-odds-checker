package cache

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Set(ctx context.Context, key, value string) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockStore) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(string), args.Error(1)
}
