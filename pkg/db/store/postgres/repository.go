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

func (r *repository) Get(obj interface{}, query string, flagScanAllOrOneRow bool, args ...interface{}) error {

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

	if flagScanAllOrOneRow {
		if err = pgxscan.ScanAll(obj, rows); err != nil {
			return fmt.Errorf("scanAll failed: %w", err)
		}
	} else {
		if err = pgxscan.ScanOne(obj, rows); err != nil {
			return fmt.Errorf("scanOne failed: %w", err)
		}
	}

	return nil
}

func (r *repository) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	ctx := r.store.GetCtx()

	var res pgconn.CommandTag
	conn, err := r.store.GetExecutor()
	if err != nil {
		return res, fmt.Errorf("get executor failed: %w", err)
	}

	res, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return res, fmt.Errorf("exec failed: %w", err)
	}

	return res, nil
}

func (r *repository) InsertOne(returnValue interface{}, query string, args ...interface{}) error {

	ctx := r.store.GetCtx()
	conn, err := r.store.GetExecutor()
	if err != nil {
		return fmt.Errorf("get executor failed: %w", err)
	}

	if returnValue != nil {
		row := conn.QueryRow(ctx, query, args...)
		if err := row.Scan(returnValue); err != nil {
			return fmt.Errorf("insertOne: QueryRow failed: %w", err)
		}
	} else {
		if _, err := conn.Exec(ctx, query, args...); err != nil {
			return fmt.Errorf("insertOne: Exec failed: %w", err)
		}
	}

	return nil
}
