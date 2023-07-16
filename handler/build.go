//go:build wireinject

package handler

import (
	"context"

	"github.com/google/wire"
)

func BuildHandler(ctx context.Context) (Handler, error) {
	panic(wire.Build(HandlerGraph))
}
