package bar

import (
	"context"

	"github.com/LeXFReeMan/operatorkit/v7/pkg/handler"
)

type Handler struct{}

func (h *Handler) EnsureCreated(ctx context.Context, req handler.Request) (*handler.Response, error) {
	return nil, nil
}

func (h *Handler) EnsureDeleted(ctx context.Context, req handler.Request) (*handler.Response, error) {
	return nil, nil
}
