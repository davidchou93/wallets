package config

import "github.com/google/wire"

var WireSet = wire.NewSet(GetConfig, GetAppEnv)
