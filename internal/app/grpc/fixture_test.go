package grpc_test

import (
	"context"
	"errors"
	gr "github.com/statistico/statistico-odds-checker/internal/app/grpc"
	"github.com/statistico/statistico-proto/statistico-data/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"testing"
)

func TestFixtureClient_Search(t *testing.T) {
	t.Run("calls fixture proto client and returns a slice of fixture struct", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockFixtureProtoClient)
		client := gr.NewFixtureClient(pc)

		request := statisticoproto.FixtureSearchRequest{}

		stream := new(gr.FixtureStream)
		ctx := context.Background()

		pc.On("Search", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)

		stream.On("Recv").Twice().Return(&statisticoproto.Fixture{}, nil)
		stream.On("Recv").Once().Return(&statisticoproto.Fixture{}, io.EOF)

		fixtures, err := client.Search(ctx, &request)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(fixtures))
		pc.AssertExpectations(t)
	})

	t.Run("returns server error if internal error returned by fixture proto client", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockFixtureProtoClient)
		client := gr.NewFixtureClient(pc)

		request := statisticoproto.FixtureSearchRequest{}

		stream := new(gr.FixtureStream)
		ctx := context.Background()

		e := status.Error(codes.Internal, "incorrect format")

		pc.On("Search", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.Search(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "Internal server error: rpc error: code = Internal desc = incorrect format", err.Error())
		pc.AssertExpectations(t)
	})

	t.Run("returns bad gateway error if non internal error returned by fixture proto client", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockFixtureProtoClient)
		client := gr.NewFixtureClient(pc)

		request := statisticoproto.FixtureSearchRequest{}

		stream := new(gr.FixtureStream)
		ctx := context.Background()

		e := errors.New("error occurred")

		pc.On("Search", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.Search(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "Bad gateway error: error occurred", err.Error())
		pc.AssertExpectations(t)
	})
}
