package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	paws "github.com/statistico/statistico-odds-checker/internal/app/publish/aws"
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
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
			ms := "{\"id\":\"1.23712\",\"eventId\":129817121,\"competitionId\":8,\"seasonId\":17420," +
				"\"round\":5,\"eventDate\":1604430059,\"name\":\"1X2\"," +
				"\"exchange\":\"BETFAIR\",\"runners\":[{\"id\":14571761,\"name\":\"Over 2.5 Goals\"," +
				"\"backPrices\":[{\"price\":1.95,\"size\":1461}],\"layPrices\":[{\"price\":1.95,\"size\":1461}]}],\"timestamp\":1604430059}"

			assert.Equal(t, ms, *i.Message)
			assert.Equal(t, "my-topic-arn", *i.TopicArn)
			return true
		})

		client.On("Publish", input).Return(&sns.PublishOutput{MessageId: aws.String("123")}, nil)

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
			ms := "{\"id\":\"1.23712\",\"eventId\":129817121,\"competitionId\":8,\"seasonId\":17420," +
				"\"round\":5,\"eventDate\":1604430059,\"name\":\"1X2\"," +
				"\"exchange\":\"BETFAIR\",\"runners\":[{\"id\":14571761,\"name\":\"Over 2.5 Goals\"," +
				"\"backPrices\":[{\"price\":1.95,\"size\":1461}],\"layPrices\":[{\"price\":1.95,\"size\":1461}]}],\"timestamp\":1604430059}"

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

func eventMarket() *stream.EventMarket {
	return &stream.EventMarket{
		ID:            "1.23712",
		EventID:       129817121,
		CompetitionID: 8,
		SeasonID:      17420,
		Round:         5,
		EventDate:     1604430059,
		MarketName:    "1X2",
		Exchange:      "BETFAIR",
		Runners: []*exchange.Runner{
			{
				ID:   14571761,
				Name: "Over 2.5 Goals",
				BackPrices: []exchange.PriceSize{
					{
						Price: 1.95,
						Size:  1461.00,
					},
				},
				LayPrices: []exchange.PriceSize{
					{
						Price: 1.95,
						Size:  1461.00,
					},
				},
			},
		},
		Timestamp: 1604430059,
	}
}
