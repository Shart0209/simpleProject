package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"simpleProject/internal/api"
)

func NewClient(ctx context.Context, cfg *api.Config) (*pgxpool.Pool, error) {
	//  Example URL
	//	postgres://jack:secret@pg.example.com:5432/mydb
	dst := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%v&connect_timeout=%d",
		cfg.Postgres.User,
		cfg.Postgres.Pswd,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DbName,
		cfg.Postgres.MaxConns,
		cfg.Postgres.ConnectTimeout,
	)

	dbpool, err := pgxpool.New(ctx, dst)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
