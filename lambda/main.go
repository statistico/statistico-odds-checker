package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf(
		"Hello %s! Your publisher is %s and your football data host is %s",
		name.Name, os.Getenv("PUBLISHER"),
		os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_HOST"),
	), nil
}

func main() {
	lambda.Start(HandleRequest)
}
