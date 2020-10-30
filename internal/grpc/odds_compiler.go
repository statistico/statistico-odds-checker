package grpc

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
)

type OddsCompilerClient interface {
	GetEventMarket(ctx context.Context, eventId uint64, market string) ([]*Odds, error)
}

type oddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

func (o *oddsCompilerClient) GetEventMarket(ctx context.Context, eventId uint64, market string) ([]*Odds, error) {
	var odds []*Odds

	req := &proto.EventRequest{
		EventId: eventId,
		Market:  market,
	}

	response, err := o.client.GetEventMarket(ctx, req)

	if err != nil {
		return odds, handleErrorResponse(err)
	}

	for _, o := range response.Odds {
		odd := &Odds{
			Price:     o.Price,
			Selection: o.Selection,
		}

		odds = append(odds, odd)
	}

	return odds, nil
}

func NewOddsCompilerClient(c proto.OddsCompilerServiceClient) OddsCompilerClient {
	return &oddsCompilerClient{client: c}
}
