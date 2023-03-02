package store

import "github.com/jackc/pgx/v5/pgconn"

type Repository interface {
	Get(obj interface{}, query string, flag bool, args ...interface{}) error
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
}
