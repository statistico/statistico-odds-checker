package grpc

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"io"
)

type FixtureClient interface {
	Search(ctx context.Context, req *proto.FixtureSearchRequest) ([]*proto.Fixture, error)
}

type fixtureClient struct {
	client proto.FixtureServiceClient
}

func (f *fixtureClient) Search(ctx context.Context, req *proto.FixtureSearchRequest) ([]*proto.Fixture, error) {
	fixtures := []*proto.Fixture{}

	stream, err := f.client.Search(ctx, req)

	if err != nil {
		return fixtures, handleErrorResponse(err)
	}

	for {
		fixture, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fixtures, &errorServerError{err}
		}

		fixtures = append(fixtures, fixture)
	}

	return fixtures, nil
}

func NewFixtureClient(p proto.FixtureServiceClient) FixtureClient {
	return &fixtureClient{client: p}
}
