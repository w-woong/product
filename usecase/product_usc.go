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

func (u *productUsc) AddProduct(ctx context.Context, o dto.Product) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	e, err := conv.ToProductEntity(&o)
	if err != nil {
		return 0, err
	}
	if e.ID == "" {
		e.CreateSetID()
	}
	res, err := u.productRepo.CreateProduct(ctx, tx, e)
	if err != nil {
		return 0, err
	}

	return res, tx.Commit()
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
