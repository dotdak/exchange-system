package infrastructure

import "github.com/google/wire"

var InfraGraph = wire.NewSet(
	NewDbConfig,
	NewMysqlDb,
	NewLogger,
)
