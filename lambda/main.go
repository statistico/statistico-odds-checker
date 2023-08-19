package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/statistico/statistico-odds-checker/internal/app/bootstrap"
)

type MyEvent struct {
	Exchange string `json:"exchange"`
	Days     int    `json:"days"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println("[INFO] Building and publishing markets")

	config := bootstrap.BuildConfig()

	return fmt.Sprintf("Exchange is %s Redis host is %s and sportmonks key is %s", event.Exchange, config.RedisConfig.Host, config.SportsMonks.ApiKey), nil

	//app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	//processor := app.Processor()
	//clock := app.Clock

	//hours := 24 * event.Days

	//from := clock.Now()
	//to := clock.Now().Add(time.Hour * time.Duration(hours))

	//if err := processor.Process(ctx, from, to, event.Exchange); err != nil {
	//	return fmt.Sprintf("[ERROR] %s\n", err.Error()), err
	//}

	//return fmt.Sprintf("[INFO] Completed for exchange %s", event.Exchange), nil
}

func main() {
	lambda.Start(HandleRequest)
}
