package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	pgStore "simpleProject/pkg/db/store"
)

type repository struct {
	executor pgStore.Executor
	store    pgStore.Store
}

func NewRepository(ex pgStore.Executor, s pgStore.Store) pgStore.Repository {
	return &repository{
		executor: ex,
		store:    s,
	}
}

func (r *repository) Get(obj interface{}, query string, flag bool, args ...interface{}) error {
	ctx := r.store.GetCtx()
	conn, err := r.store.GetExecutor()
	if err != nil {
		return err
	}

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if flag {
		if err = scanMany(rows, obj); err != nil {
			return fmt.Errorf("getAll: scanAll failed: %w", err)
		}
	} else {
		if err = scanOne(rows, obj); err != nil {
			return fmt.Errorf("getOne: scanOne failed: %w", err)
		}
	}

	return nil
}

func (r *repository) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	ctx := r.store.GetCtx()

	var res pgconn.CommandTag
	ex, err := r.store.GetExecutor()
	if err != nil {
		return res, fmt.Errorf("get executor failed: %w", err)
	}

	res, err = ex.Exec(ctx, query, args...)
	if err != nil {
		return res, fmt.Errorf("exec failed: %w", err)
	}

	return res, nil
}
