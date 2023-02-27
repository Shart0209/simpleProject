package postgres

import (
	"fmt"
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

func (r *repository) GetAll() error {
	//TODO implement me
	panic("implement me")

}

func (r *repository) GetByName(obj interface{}, query string, args ...interface{}) error {
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

	for rows.Next() {
		if err = rows.Scan(&obj.contract_id, &description); err != nil {

		}
	}

	fmt.Println(rows)

	return nil
}

func (r *repository) Create() error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
