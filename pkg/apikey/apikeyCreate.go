package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// createResp describes the structure of the JSON document returned by the API.
type createResp struct {
	Token string `json:"token"`
}

// Create creates a new API key and returns the value of newly created API key
// and the raw API response.
func (a *apiKey) Create() (string, string, error) {
	client, err := _client.New(ClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(a.endpoint, a.orgSlug)

	var data []_client.Data
	data = append(data, _client.Data{Key: "name", Value: a.name})
	data = append(data, _client.Data{Key: "role", Value: a.role})

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return "", "", errors.New("API key with this name already exists")
		} else {
			return "", "", err
		}
	}

	var jsonData createResp
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return jsonData.Token, string(body), nil
}