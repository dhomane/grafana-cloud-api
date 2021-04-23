package grafana

import (
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Restart restarts a Stack and returns the raw API response.
func (g *Grafana) Restart() (string, error) {
	client, err := _client.New(g.ClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(g.Endpoint+"/restart", g.StackSlug)

	body, statusCode, err := client.Post(nil)
	if err != nil {
		if statusCode == 404 {
			return "", fmt.Errorf("Stack Slug not found: %s", err)
		}

		return "", err
	}

	return string(body), nil
}
