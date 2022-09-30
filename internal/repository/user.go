package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/davidchou93/wallets/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepo RepositoryImpl

func (s *UserRepo) GetList(ctx context.Context, option model.GetUserOption) ([]model.User, error) {
	var (
		builder squirrel.SelectBuilder
		query   string
		args    []interface{}
		err     error
		result  []model.User
	)
	builder = s.sq.Select("*").From("users")
	limit := model.DBDefaultLimit
	if option.Limit.Valid {
		limit = uint64(option.Limit.Int64)
	}
	builder = builder.Limit(limit)

	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}
	err = s.db.SelectContext(ctx, &result, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return result, err
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return (*UserRepo)(NewRepository(db))
}
