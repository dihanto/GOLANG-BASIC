package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dihanto/golang-basic/17-database-mysql/entity"
)

type productUsecase struct {
	productRepo    entity.ProductRepository
	contextTimeout time.Duration
}

func NewProductUsecase(pr entity.ProductRepository, timeout time.Duration) entity.ProductUsecase {
	return &productUsecase{
		productRepo:    pr,
		contextTimeout: timeout,
	}
}
func (pu *productUsecase) Insert(c context.Context, product *entity.Product) (err error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Insert(ctx, product); err != nil {
		return
	}
	fmt.Println(product)
	return
}

func (pu *productUsecase) FindAll(c context.Context) ([]entity.Product, error) {
	panic("not implemented") // TODO: Implement
}
