package store

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
)

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Repository interface {
	Get(obj interface{}, query string, flagScanAllOrOne bool, args ...interface{}) error
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
	InsertOne(returnValue interface{}, query string, args ...interface{}) error
}

type Store interface {
	Stop() error
	GetLogger() zerolog.Logger
	GetExecutor() (Executor, error)

	GetRepository() Repository
}

type BaseStore interface {
	GetCtxWithTimeout() (context.Context, context.CancelFunc)
	GetExecutor() (Executor, error)
}
