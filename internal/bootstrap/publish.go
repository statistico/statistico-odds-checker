package bootstrap

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/statistico/statistico-odds-checker/internal/publish"
	pa "github.com/statistico/statistico-odds-checker/internal/publish/aws"
	"github.com/statistico/statistico-odds-checker/internal/publish/log"
)

func (c Container) Publisher() publish.Publisher {
	if c.Config.Publisher == "aws" {
		key := c.Config.AwsConfig.Key
		secret := c.Config.AwsConfig.Secret

		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
			Region:      aws.String(c.Config.AwsConfig.Region),
		})

		if err != nil {
			panic(err)
		}

		return pa.NewPublisher(sns.New(sess), c.Config.AwsConfig.TopicArn)
	}

	if c.Config.Publisher == "log" {
		return log.NewPublisher(c.Logger)
	}

	panic("Queue driver provided is not supported")
}
