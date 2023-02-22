package postgres

import (
	pgStore "simpleProject/pkg/store/client"
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

func (r *repository) All() error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetByName(id int64) error {
	//TODO implement me
	panic("implement me")
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
