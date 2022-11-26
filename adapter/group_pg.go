package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/product/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type groupPg struct {
	db *gorm.DB
}

func NewGroupPg(db *gorm.DB) *groupPg {
	return &groupPg{
		db: db,
	}
}

func (a *groupPg) CreateGroup(ctx context.Context, tx common.TxController, group entity.Group) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&group)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}

func (a *groupPg) ReadGroup(ctx context.Context, tx common.TxController, id string) (entity.Group, error) {
	return a.readGroup(ctx, tx.(*txcom.GormTxController).Tx, id)
}

func (a *groupPg) ReadGroupNoTx(ctx context.Context, id string) (entity.Group, error) {
	return a.readGroup(ctx, a.db, id)
}

func (a *groupPg) readGroup(ctx context.Context, db *gorm.DB, id string) (entity.Group, error) {
	var group entity.Group
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Preload(clause.Associations).
		Limit(1).Find(&group)
	if res.Error != nil {
		return entity.NilGroup, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilGroup, common.ErrRecordNotFound
	}

	return group, nil
}

func (a *groupPg) UpdateGroup(ctx context.Context, tx common.TxController, id string, group entity.Group) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Model(&group).
		Where("id = ?", id).
		Updates(&group)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *groupPg) DeleteGroup(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.Group{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil

}
