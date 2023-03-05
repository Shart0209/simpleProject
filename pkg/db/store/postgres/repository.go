package postgres

import (
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
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
		return fmt.Errorf("get executor failed: %w", err)
	}

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if flag {
		if err = pgxscan.ScanAll(obj, rows); err != nil {
			return fmt.Errorf("scanAll failed: %w", err)
		}
	} else {
		if err = pgxscan.ScanRow(obj, rows); err != nil {
			return fmt.Errorf("scanRow failed: %w", err)
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

//func (r *repository) InsertOne() error {
//	ctx := r.store.GetCtx()
//	conn, err := r.store.GetExecutor()
//	if err != nil {
//		return fmt.Errorf("get executor failed: %w", err)
//	}
//
//	return nil
//}
