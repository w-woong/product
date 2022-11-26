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

type GroupRepo interface {
	CreateGroup(ctx context.Context, tx common.TxController, group entity.Group) (int64, error)
	ReadGroup(ctx context.Context, tx common.TxController, id string) (entity.Group, error)
	ReadGroupNoTx(ctx context.Context, id string) (entity.Group, error)
	UpdateGroup(ctx context.Context, tx common.TxController, id string, group entity.Group) (int64, error)
	DeleteGroup(ctx context.Context, tx common.TxController, id string) (int64, error)
}

type GroupProductRepo interface {
	CreateGroupProduct(ctx context.Context, tx common.TxController, groupProduct entity.GroupProduct) (int64, error)
	ReadGroupProduct(ctx context.Context, tx common.TxController, groupID, productID string) (entity.GroupProduct, error)
	ReadGroupProductNoTx(ctx context.Context, groupID, productID string) (entity.GroupProduct, error)
	UpdateGroupProduct(ctx context.Context, tx common.TxController, groupID, productID string, groupProduct entity.GroupProduct) (int64, error)
	DeleteGroupProduct(ctx context.Context, tx common.TxController, groupID, productID string) (int64, error)
	DeleteByGroupID(ctx context.Context, tx common.TxController, groupID string) (int64, error)
}
