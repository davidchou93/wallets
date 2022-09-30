package logging

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewLogger,
)
