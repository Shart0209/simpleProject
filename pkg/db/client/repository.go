package client

import "context"

type Repository interface {
	Create(ctx context.Context) error
	All(ctx context.Context) error
	GetByName(ctx context.Context, name string) error
	Update(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}
