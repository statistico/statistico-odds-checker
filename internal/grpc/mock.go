package grpc

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type FixtureStream struct {
	mock.Mock
	grpc.ClientStream
}

func (f *FixtureStream) Recv() (*proto.Fixture, error) {
	args := f.Called()
	return args.Get(0).(*proto.Fixture), args.Error(1)
}

type MockFixtureGrpcClient struct {
	mock.Mock
}

func (m *MockFixtureGrpcClient) Search(ctx context.Context, req *proto.FixtureSearchRequest) ([]*proto.Fixture, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*proto.Fixture), args.Error(1)
}

type MockFixtureProtoClient struct {
	mock.Mock
}

func (f *MockFixtureProtoClient) ListSeasonFixtures(ctx context.Context, in *proto.SeasonFixtureRequest, opts ...grpc.CallOption) (proto.FixtureService_ListSeasonFixturesClient, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(proto.FixtureService_ListSeasonFixturesClient), args.Error(1)
}

func (f *MockFixtureProtoClient) FixtureByID(ctx context.Context, in *proto.FixtureRequest, opts ...grpc.CallOption) (*proto.Fixture, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(*proto.Fixture), args.Error(1)
}

func (f *MockFixtureProtoClient) Search(ctx context.Context, in *proto.FixtureSearchRequest, opts ...grpc.CallOption) (proto.FixtureService_SearchClient, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(proto.FixtureService_SearchClient), args.Error(1)
}
