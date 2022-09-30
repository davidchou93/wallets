package model

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	WithTransaction(ctx context.Context, fn TxFunc) error
	UserRepo() UserRepo
}

// TxFunc is the func for WithTransaction call
type TxFunc func(tx *sqlx.Tx) error

const DBDefaultLimit uint64 = 500
