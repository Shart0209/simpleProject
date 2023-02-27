package store

type Repository interface {
	GetAll() error
	GetByName(obj interface{}, query string, args ...interface{}) error
	Create() error
	Update(id int64) error
	Delete(id int64) error
}
