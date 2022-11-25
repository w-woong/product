package port

import (
	"context"

	"github.com/w-woong/product/dto"
)

type ProductUsc interface {
	FindProduct(ctx context.Context, id string) (dto.Product, error)
}