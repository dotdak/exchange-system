package handler

import (
	"github.com/dotdak/exchange-system/infrastructure"
	"github.com/dotdak/exchange-system/repo"
	"github.com/google/wire"
)

var HandlerGraph = wire.NewSet(
	infrastructure.InfraGraph,
	repo.RepoGraph,
	NewHandler,
)
