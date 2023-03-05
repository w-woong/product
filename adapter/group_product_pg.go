package adapter

import (
	"context"

	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/product/entity"
	"github.com/w-woong/product/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ port.GroupProductRepo = (*groupProductPg)(nil)

type groupProductPg struct {
	db *gorm.DB
}

func NewGroupProductPg(db *gorm.DB) *groupProductPg {
	return &groupProductPg{
		db: db,
	}
}

func (a *groupProductPg) CreateGroupProduct(ctx context.Context, tx commonport.TxController,
	groupProduct entity.GroupProduct) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).Create(&groupProduct)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *groupProductPg) ReadGroupProduct(ctx context.Context, tx commonport.TxController,
	groupID, productID string) (entity.GroupProduct, error) {
	return a.readGroupProduct(ctx, tx.(*txcom.GormTxController).Tx, groupID, productID)
}

func (a *groupProductPg) ReadGroupProductNoTx(ctx context.Context,
	groupID, productID string) (entity.GroupProduct, error) {
	return a.readGroupProduct(ctx, a.db, groupID, productID)
}

func (a *groupProductPg) readGroupProduct(ctx context.Context, db *gorm.DB,
	groupID, productID string) (entity.GroupProduct, error) {

	var groupProduct entity.GroupProduct
	res := db.WithContext(ctx).Where("group_id = ? and product_id = ?", groupID, productID).
		Preload(clause.Associations).Limit(1).Find(&groupProduct)
	if res.Error != nil {
		return entity.NilGroupProduct, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilGroupProduct, commondto.ErrRecordNotFound
	}

	return groupProduct, nil
}

func (a *groupProductPg) UpdateGroupProduct(ctx context.Context, tx commonport.TxController,
	groupID, productID string, groupProduct entity.GroupProduct) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&groupProduct).
		Where("group_id = ? and product_id = ?", groupID, productID).
		Updates(&groupProduct)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *groupProductPg) DeleteGroupProduct(ctx context.Context, tx commonport.TxController, groupID, productID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("group_id = ? and product_id = ?", groupID, productID).
		Delete(&entity.GroupProduct{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
func (a *groupProductPg) DeleteByGroupID(ctx context.Context, tx commonport.TxController,
	groupID string) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("group_id = ?", groupID).
		Delete(&entity.GroupProduct{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
func (a *groupProductPg) DeleteByProductID(ctx context.Context, tx commonport.TxController, productID string) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("product_id = ?", productID).
		Delete(&entity.GroupProduct{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}
