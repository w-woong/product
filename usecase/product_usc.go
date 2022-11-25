package usecase

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/product/conv"
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/port"
)

type productUsc struct {
	beginner common.TxBeginner
	repo     port.ProductRepo
}

func NewProductUsc(beginner common.TxBeginner, repo port.ProductRepo) *productUsc {
	return &productUsc{
		beginner: beginner,
		repo:     repo,
	}
}

func (u *productUsc) FindProduct(ctx context.Context, id string) (dto.Product, error) {
	product, err := u.repo.ReadProductNoTx(ctx, id)
	if err != nil {
		return dto.NilProduct, err
	}

	return conv.ToProductDto(&product)
}
