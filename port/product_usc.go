package port

import (
	"context"

	"github.com/w-woong/product/dto"
)

type ProductUsc interface {
	AddProduct(ctx context.Context, o dto.Product) (int64, error)
	FindProduct(ctx context.Context, id string) (dto.Product, error)
	FindProductsByGroupID(ctx context.Context, groupID string) (dto.ProductList, error)
}

type GroupUsc interface {
	AddGroup(ctx context.Context, o dto.Group) (int64, error)
	FindGroup(ctx context.Context, id string) (dto.Group, error)
}
