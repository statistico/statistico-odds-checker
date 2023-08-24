package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ParameterStoreValue struct {
	Parameter struct {
		Value string
	}
}

func getSsmParameter(v string) string {
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

	response := ParameterStoreValue{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err.Error())
	}

	return response.Parameter.Value
}
