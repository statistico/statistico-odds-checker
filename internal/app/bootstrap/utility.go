package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type parameterStoreValue struct {
	parameter struct {
		value string
	}
}

type awsClient struct {
	httpClient *http.Client
}

func (a *awsClient) getSsmParameter(v string) string {
	url := fmt.Sprintf(
		"http://localhost:2773/systemsmanager/parameters/get/?name=%s&withDecryption=true",
		v,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("X-Aws-Parameters-Secrets-Token", os.Getenv("AWS_SESSION_TOKEN"))

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	response := parameterStoreValue{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err.Error())
	}

	fmt.Println(response.parameter.value)

	return response.parameter.value
}

func newAwsClient() *awsClient {
	return &awsClient{httpClient: &http.Client{}}
}
