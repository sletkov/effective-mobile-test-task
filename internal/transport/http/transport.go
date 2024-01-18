package httptransport

import (
	"context"
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

func (t *Transport) Get(ctx context.Context, url string) (*http.Response, error) {
	response, err := t.client.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
