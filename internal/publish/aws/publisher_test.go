package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	paws "github.com/statistico/statistico-odds-checker/internal/publish/aws"
	"github.com/statistico/statistico-odds-checker/internal/sport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPublisher_PublishMarket(t *testing.T) {
	t.Run("publishes new message via AWS SNS client", func(t *testing.T) {
		t.Helper()

		client := new(paws.MockSnsClient)
		topic := "my-topic-arn"

		publisher := paws.NewPublisher(client, topic)

		input := mock.MatchedBy(func(i *sns.PublishInput) bool {
			ms := "{\"eventId\":129817121,\"sport\":\"football\",\"eventDate\":1604430000,\"name\":\"1X2\",\"side\":\"BACK\"," +
				"\"exchange\":\"betfair\",\"exchangeMarket\":{\"id\":\"1.2817676712\",\"runners\":[{\"id\":14571761,\"name\":\"\"," +
				"\"prices\":[{\"price\":1.95,\"size\":1461}]}]},\"statisticoOdds\":[{\"price\":1.45,\"selection\":\"over\"}," +
				"{\"price\":2.78,\"selection\":\"under\"}],\"timestamp\":1604430059}"

			assert.Equal(t, ms, *i.Message)
			assert.Equal(t, "my-topic-arn", *i.TopicArn)
			return true
		})

		client.On("Publish", input).Return(&sns.PublishOutput{}, nil)

		err := publisher.PublishMarket(eventMarket())

		if err != nil {
			t.Fatalf("Expected nil, got %s", err)
		}
	})

	t.Run("returns error if returned by AWS SNS client", func(t *testing.T) {
		t.Helper()

		client := new(paws.MockSnsClient)
		topic := "my-topic-arn"

		publisher := paws.NewPublisher(client, topic)

		input := mock.MatchedBy(func(i *sns.PublishInput) bool {
			ms := "{\"eventId\":129817121,\"sport\":\"football\",\"eventDate\":1604430000,\"name\":\"1X2\",\"side\":\"BACK\"," +
				"\"exchange\":\"betfair\",\"exchangeMarket\":{\"id\":\"1.2817676712\",\"runners\":[{\"id\":14571761,\"name\":\"\"," +
				"\"prices\":[{\"price\":1.95,\"size\":1461}]}]},\"statisticoOdds\":[{\"price\":1.45,\"selection\":\"over\"}," +
				"{\"price\":2.78,\"selection\":\"under\"}],\"timestamp\":1604430059}"

			assert.Equal(t, ms, *i.Message)
			assert.Equal(t, "my-topic-arn", *i.TopicArn)
			return true
		})

		client.On("Publish", input).Return(&sns.PublishOutput{}, errors.New("error"))

		err := publisher.PublishMarket(eventMarket())

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func eventMarket() *sport.EventMarket {
	return &sport.EventMarket{
		EventID:        129817121,
		Sport:          "football",
		EventDate:      1604430000,
		MarketName:     "1X2",
		Side:           "BACK",
		Exchange:       "betfair",
		ExchangeMarket: exchange.Market{
			ID:      "1.2817676712",
			Runners: []exchange.Runner{
				{
					ID: 14571761,
					Prices: []exchange.PriceSize{
						{
							Price: 1.95,
							Size: 1461.00,
						},
					},
				},
			},
		},
		StatisticoOdds: []*proto.Odds{
			{
				Price: 1.45,
				Selection: "over",
			},
			{
				Price: 2.78,
				Selection: "under",
			},
		},
		Timestamp:      1604430059,
	}
}