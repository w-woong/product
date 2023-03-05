package adapter

import (
	"context"

	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/product/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productPg struct {
	db *gorm.DB
}

func NewProductPg(db *gorm.DB) *productPg {
	return &productPg{
		db: db,
	}
}

func (a *productPg) CreateProduct(ctx context.Context, tx commonport.TxController, product entity.Product) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&product)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *productPg) ReadProduct(ctx context.Context, tx commonport.TxController, id string) (entity.Product, error) {

	return a.readProduct(ctx, tx.(*txcom.GormTxController).Tx, id)
}

func (a *productPg) ReadProductNoTx(ctx context.Context, id string) (entity.Product, error) {
	return a.readProduct(ctx, a.db, id)
}

func (a *productPg) readProduct(ctx context.Context, db *gorm.DB, id string) (entity.Product, error) {

	var product entity.Product
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Preload(clause.Associations).
		Limit(1).Find(&product)
	if res.Error != nil {
		return entity.NilProduct, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilProduct, commondto.ErrRecordNotFound
	}
	return product, nil
}

func (a *productPg) UpdateProduct(ctx context.Context, tx commonport.TxController, id string, product entity.Product) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&product).
		Where("id = ?", id).
		Updates(&product)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *productPg) DeleteProduct(ctx context.Context, tx commonport.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.Product{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
