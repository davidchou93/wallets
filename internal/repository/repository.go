package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/davidchou93/wallets/internal/model"
	"github.com/jmoiron/sqlx"
)

type RepositoryImpl struct {
	db *sqlx.DB
	sq squirrel.StatementBuilderType
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	repo := new(RepositoryImpl)
	repo.db = db
	repo.sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return repo
}

func (r *RepositoryImpl) UserRepo() model.UserRepo {
	return (*UserRepo)(r)
}
func (r *RepositoryImpl) WithTransaction(ctx context.Context, fn model.TxFunc) error {
	return withTransaction(ctx, r.db, TxFn(fn))
}

// TxFn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(tx *sqlx.Tx) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func withTransaction(ctx context.Context, db *sqlx.DB, fn TxFn) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
