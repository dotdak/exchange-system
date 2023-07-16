package repo

import "github.com/google/wire"

var RepoGraph = wire.NewSet(
	NewWagerRepo,
	NewBuyRepo,
)
