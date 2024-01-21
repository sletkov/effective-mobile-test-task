package httptransport

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

type Transport struct {
	client *http.Client
}

func New(client *http.Client) *Transport {
	return &Transport{
		client: client,
	}
}

// Make GET request by url
func (t *Transport) Get(ctx context.Context, url string) (*http.Response, error) {
	slog.DebugContext(ctx, fmt.Sprintf("making GET request to %s", url))

	response, err := t.client.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
