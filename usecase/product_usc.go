package usecase

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/product/conv"
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/port"
)

type productUsc struct {
	beginner    common.TxBeginner
	productRepo port.ProductRepo
	groupRepo   port.GroupRepo
}

func NewProductUsc(beginner common.TxBeginner,
	productRepo port.ProductRepo, groupRepo port.GroupRepo) *productUsc {

	return &productUsc{
		beginner:    beginner,
		productRepo: productRepo,
		groupRepo:   groupRepo,
	}
}

func (u *productUsc) FindProduct(ctx context.Context, id string) (dto.Product, error) {
	product, err := u.productRepo.ReadProductNoTx(ctx, id)
	if err != nil {
		return dto.NilProduct, err
	}

	return conv.ToProductDto(&product)
}

func (u *productUsc) FindProductsByGroupID(ctx context.Context,
	groupID string) (dto.ProductList, error) {

	group, err := u.groupRepo.ReadGroupNoTx(ctx, groupID)
	if err != nil {
		return nil, err
	}
	groupDto, err := conv.ToGroupDto(&group)
	if err != nil {
		return nil, err
	}
	return groupDto.Products, nil
}
