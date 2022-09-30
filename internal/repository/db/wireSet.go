package db

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewDBConfig,
	NewPgConn,
)
