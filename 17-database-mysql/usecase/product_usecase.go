package usecase

import (
	"context"
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
	return
}

func (pu *productUsecase) FindAll(c context.Context) (products []entity.Product, err error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	products, err = pu.productRepo.FindAll(ctx)
	if err != nil {
		return
	}
	return
}

func (pu *productUsecase) Update(ctx context.Context, product *entity.Product) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Update(ctx, product); err != nil {
		return
	}
	return
}

func (pu *productUsecase) Delete(ctx context.Context, id int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Delete(ctx, id); err != nil {
		return
	}
	return
}
