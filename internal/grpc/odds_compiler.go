package grpc

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
)

type OddsCompilerClient interface {
	GetEventMarket(ctx context.Context, eventId uint64, market string) ([]*proto.Odds, error)
}

type oddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

func (o *oddsCompilerClient) GetEventMarket(ctx context.Context, eventId uint64, market string) ([]*proto.Odds, error) {
	req := &proto.EventRequest{
		EventId: eventId,
		Market:  market,
	}

	response, err := o.client.GetEventMarket(ctx, req)

	if err != nil {
		return []*proto.Odds{}, handleErrorResponse(err)
	}

	return response.Odds, nil
}

func NewOddsCompilerClient(c proto.OddsCompilerServiceClient) OddsCompilerClient {
	return &oddsCompilerClient{client: c}
}
