package httptransport

import (
	"context"
	"net/http"
)

const URL = ""

type Transport struct {
	client *http.Client
}

func New(client *http.Client) *Transport {
	return &Transport{
		client: client,
	}
}

// func (t *Transport) Agify(ctx context.Context, name string) (*http.Response, error) {
// 	url := "https://api.agify.io/?name=" + name

// 	response, err := t.client.Get(url)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (t *Transport) Genderize(ctx context.Context, name string) (*http.Response, error) {
// 	url := "https://api.genderize.io/?name=" + name

// 	response, err := t.client.Get(url)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (t *Transport) Nationalize(ctx context.Context, name string) (*http.Response, error) {
// 	url := "https://api.nationalize.io/?name=" + name

// 	response, err := t.client.Get(url)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }

func (t *Transport) Get(ctx context.Context, url string) (*http.Response, error) {
	response, err := t.client.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
