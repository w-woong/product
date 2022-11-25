package port

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/product/entity"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, tx common.TxController, product entity.Product) (int64, error)
	ReadProduct(ctx context.Context, tx common.TxController, id string) (entity.Product, error)
	ReadProductNoTx(ctx context.Context, id string) (entity.Product, error)
	UpdateProduct(ctx context.Context, tx common.TxController, id string, product entity.Product) (int64, error)
	DeleteProduct(ctx context.Context, tx common.TxController, id string) (int64, error)
}
