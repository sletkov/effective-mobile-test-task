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
	slog.InfoContext(ctx, fmt.Sprintf("transport: making GET request to %s", url))

	response, err := t.client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("transport: making get request to %s: %w", url, err)
	}

	slog.DebugContext(ctx, fmt.Sprintf("transport: got response %v", response))

	slog.InfoContext(ctx, "transport: request was made successfully")

	return response, nil
}
