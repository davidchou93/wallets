package model

import (
	"context"
	"database/sql"
)

type User struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UserService interface {
	GetList(ctx context.Context, option GetUserOption) ([]User, error)
}

type GetUserOption struct {
	Limit sql.NullInt64
}

type UserRepo interface {
	GetList(ctx context.Context, option GetUserOption) ([]User, error)
}
