package grpc

import (
	"context"
	"github.com/statistico/statistico-proto/statistico-data/go"
	"io"
)

type FixtureClient interface {
	Search(ctx context.Context, req *statisticoproto.FixtureSearchRequest) ([]*statisticoproto.Fixture, error)
}

type fixtureClient struct {
	client statisticoproto.FixtureServiceClient
}

func (f *fixtureClient) Search(ctx context.Context, req *statisticoproto.FixtureSearchRequest) ([]*statisticoproto.Fixture, error) {
	fixtures := []*statisticoproto.Fixture{}

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

func NewFixtureClient(p statisticoproto.FixtureServiceClient) FixtureClient {
	return &fixtureClient{client: p}
}
