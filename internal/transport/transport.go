package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	// Agify(ctx context.Context) error
	// Genderize(ctx context.Context) error
	// Nationalize(ctx context.Context) error
	Get(ctx context.Context, url string) (*http.Response, error)
}
