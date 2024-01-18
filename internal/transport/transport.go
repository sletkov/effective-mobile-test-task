package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Get(ctx context.Context, url string) (*http.Response, error)
}
