package entity

import "context"

type Product struct {
	ID    int64
	Name  string
	Price float32
}
type ProductRepository interface {
	Insert(ctx context.Context, product *Product) error
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int) error
}
type ProductUsecase interface {
	Insert(ctx context.Context, product *Product) error
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int) error
}
