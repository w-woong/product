package port

import (
	"context"

	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/product/entity"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, tx commonport.TxController, product entity.Product) (int64, error)
	ReadProduct(ctx context.Context, tx commonport.TxController, id string) (entity.Product, error)
	ReadProductNoTx(ctx context.Context, id string) (entity.Product, error)
	UpdateProduct(ctx context.Context, tx commonport.TxController, id string, product entity.Product) (int64, error)
	DeleteProduct(ctx context.Context, tx commonport.TxController, id string) (int64, error)
}

type GroupRepo interface {
	CreateGroup(ctx context.Context, tx commonport.TxController, group entity.Group) (int64, error)
	ReadGroup(ctx context.Context, tx commonport.TxController, id string) (entity.Group, error)
	ReadGroupNoTx(ctx context.Context, id string) (entity.Group, error)
	UpdateGroup(ctx context.Context, tx commonport.TxController, id string, group entity.Group) (int64, error)
	DeleteGroup(ctx context.Context, tx commonport.TxController, id string) (int64, error)
}

type GroupProductRepo interface {
	CreateGroupProduct(ctx context.Context, tx commonport.TxController, groupProduct entity.GroupProduct) (int64, error)
	ReadGroupProduct(ctx context.Context, tx commonport.TxController, groupID, productID string) (entity.GroupProduct, error)
	ReadGroupProductNoTx(ctx context.Context, groupID, productID string) (entity.GroupProduct, error)
	UpdateGroupProduct(ctx context.Context, tx commonport.TxController, groupID, productID string, groupProduct entity.GroupProduct) (int64, error)
	DeleteGroupProduct(ctx context.Context, tx commonport.TxController, groupID, productID string) (int64, error)
	DeleteByGroupID(ctx context.Context, tx commonport.TxController, groupID string) (int64, error)
	DeleteByProductID(ctx context.Context, tx commonport.TxController, productID string) (int64, error)
}
