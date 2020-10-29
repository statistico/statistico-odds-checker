package grpc_test

import (
	"context"
	"errors"
	gr "github.com/statistico/statistico-odds-checker/internal/grpc"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestOddsCompilerClient_GetEventMarket(t *testing.T) {
	t.Run("calls odds compiler proto client and returns a slice of proto odds struct", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockOddsCompilerProtoClient)
		client := gr.NewOddsCompilerClient(pc)

		market := proto.EventMarket{
			EventId: 38192,
			Market:  "OVER_UNDER_25",
			Odds: []*proto.Odds{
				{
					Price:     1.45,
					Selection: "over",
				},
				{
					Price:     3.45,
					Selection: "under",
				},
			},
		}

		req := mock.MatchedBy(func(r *proto.EventRequest) bool {
			assert.Equal(t, uint64(38192), r.EventId)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		ctx := context.Background()

		pc.On("GetEventMarket", ctx, req, []grpc.CallOption(nil)).Return(&market, nil)

		odds, err := client.GetEventMarket(ctx, uint64(38192), "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, market.Odds, odds)
		pc.AssertExpectations(t)
	})

	t.Run("returns server error if internal error returns by odds compiler proto client", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockOddsCompilerProtoClient)
		client := gr.NewOddsCompilerClient(pc)

		req := mock.MatchedBy(func(r *proto.EventRequest) bool {
			assert.Equal(t, uint64(38192), r.EventId)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		ctx := context.Background()

		e := status.Error(codes.Internal, "incorrect format")

		pc.On("GetEventMarket", ctx, req, []grpc.CallOption(nil)).Return(&proto.EventMarket{}, e)

		_, err := client.GetEventMarket(ctx, uint64(38192), "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Error(t, err, gr.ErrorServerError)
		pc.AssertExpectations(t)
	})

	t.Run("returns bad gateway error if non internal error returned by odds compiler proto client", func(t *testing.T) {
		t.Helper()

		pc := new(gr.MockOddsCompilerProtoClient)
		client := gr.NewOddsCompilerClient(pc)

		req := mock.MatchedBy(func(r *proto.EventRequest) bool {
			assert.Equal(t, uint64(38192), r.EventId)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		ctx := context.Background()

		e := errors.New("error occurred")

		pc.On("GetEventMarket", ctx, req, []grpc.CallOption(nil)).Return(&proto.EventMarket{}, e)

		_, err := client.GetEventMarket(ctx, uint64(38192), "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Error(t, err, gr.ErrorBadGateway)
		pc.AssertExpectations(t)
	})
}
