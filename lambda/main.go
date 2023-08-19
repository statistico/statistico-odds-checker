package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
)

type MyEvent struct {
	Name string `json:"name"`
}

type ParameterStoreValue struct {
	Parameter struct {
		Value string
	}
}

func getSsmParameter(v string) (string, error) {
	url := fmt.Sprintf(
		"http://localhost:2773/systemsmanager/parameters/get/?name=%s&withDecryption=true",
		v,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("X-Aws-Parameters-Secrets-Token", os.Getenv("AWS_SESSION_TOKEN"))

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	response := ParameterStoreValue{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Parameter.Value, nil
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	val, err := getSsmParameter("statistico-odds-checker-BETFAIR_USERNAME")

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Hello %s! Your username is %s and your football data host is %s",
		name.Name,
		val,
		os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_HOST"),
	), nil
}

func main() {
	lambda.Start(HandleRequest)
}
