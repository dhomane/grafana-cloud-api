package stack

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// apiKey described properties of the API key returned by the API.
type apiKey struct {
	Key string `json:"key"`
}

// ApiKeyCreate creates a new Stack API key and returns the value of the newly
// created API key and the raw API response.
func (s *stack) ApiKeyCreate() (string, string, error) {
	client, err := _client.New(ClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(s.endpoint+"/%s/api/auth/keys", s.stackSlug)

	var data []_client.Data
	data = append(data, _client.Data{Key: "name", Value: s.name})
	data = append(data, _client.Data{Key: "role", Value: s.role})

	if s.secondsToLive != "" {
		data = append(data, _client.Data{Key: "secondsToLive", Value: s.secondsToLive})
	}

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return "", "", errors.New("API key with this name already exists")
		} else {
			return "", "", err
		}
	}

	var jsonData apiKey
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return jsonData.Key, string(body), nil
}
