package model

import (
	"errors"
)

var (
	ErrGeneral          = errors.New("general error")
	ErrBadRequestFormat = errors.New("request format not valid")
	ErrBadPermission    = errors.New("bad permission")
)
