package store

type Repository interface {
	Create() error
	All() error
	GetByName(id int64) error
	Update(id int64) error
	Delete(id int64) error
}
