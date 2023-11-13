// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dao

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (User, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	QueryUser(ctx context.Context, arg QueryUserParams) ([]User, int64, error)
}

var _ Querier = (*Queries)(nil)
