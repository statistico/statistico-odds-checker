package aws

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	"github.com/statistico/statistico-odds-checker/internal/app/sport"
)

type Publisher struct {
	client   snsiface.SNSAPI
	topicArn string
}

func (p *Publisher) PublishMarket(m *sport.EventMarket) error {
	js, err := json.Marshal(m)

	if err != nil {
		return err
	}

	input := sns.PublishInput{
		Message:  aws.String(string(js)),
		TopicArn: aws.String(p.topicArn),
	}

	if _, err = p.client.Publish(&input); err != nil {
		return err
	}

	return nil
}

func NewPublisher(c snsiface.SNSAPI, t string) publish.Publisher {
	return &Publisher{
		client:   c,
		topicArn: t,
	}
}
