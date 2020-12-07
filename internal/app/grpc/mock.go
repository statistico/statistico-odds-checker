package grpc

import (
	"context"
	"github.com/statistico/statistico-proto/statistico-data/go"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type FixtureStream struct {
	mock.Mock
	grpc.ClientStream
}

func (f *FixtureStream) Recv() (*statisticoproto.Fixture, error) {
	args := f.Called()
	return args.Get(0).(*statisticoproto.Fixture), args.Error(1)
}

type MockFixtureGrpcClient struct {
	mock.Mock
}

func (m *MockFixtureGrpcClient) Search(ctx context.Context, req *statisticoproto.FixtureSearchRequest) ([]*statisticoproto.Fixture, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*statisticoproto.Fixture), args.Error(1)
}

type MockFixtureProtoClient struct {
	mock.Mock
}

func (f *MockFixtureProtoClient) ListSeasonFixtures(ctx context.Context, in *statisticoproto.SeasonFixtureRequest, opts ...grpc.CallOption) (statisticoproto.FixtureService_ListSeasonFixturesClient, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(statisticoproto.FixtureService_ListSeasonFixturesClient), args.Error(1)
}

func (f *MockFixtureProtoClient) FixtureByID(ctx context.Context, in *statisticoproto.FixtureRequest, opts ...grpc.CallOption) (*statisticoproto.Fixture, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(*statisticoproto.Fixture), args.Error(1)
}

func (f *MockFixtureProtoClient) Search(ctx context.Context, in *statisticoproto.FixtureSearchRequest, opts ...grpc.CallOption) (statisticoproto.FixtureService_SearchClient, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(statisticoproto.FixtureService_SearchClient), args.Error(1)
}
