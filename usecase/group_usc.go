package usecase

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/product/conv"
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/port"
)

type groupUsc struct {
	beginner  common.TxBeginner
	groupRepo port.GroupRepo
}

func NewGroupUsc(beginner common.TxBeginner, groupRepo port.GroupRepo) *groupUsc {
	return &groupUsc{
		beginner:  beginner,
		groupRepo: groupRepo,
	}
}

func (u *groupUsc) AddGroup(ctx context.Context, o dto.Group) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	e, err := conv.ToGroupEntity(&o)
	if err != nil {
		return 0, err
	}
	if e.ID == "" {
		e.CreateSetID()
	}

	res, err := u.groupRepo.CreateGroup(ctx, tx, e)
	if err != nil {
		return 0, err
	}
	return res, tx.Commit()
}
func (u *groupUsc) FindGroup(ctx context.Context, id string) (dto.Group, error) {

	group, err := u.groupRepo.ReadGroupNoTx(ctx, id)
	if err != nil {
		return dto.NilGroup, err
	}
	return conv.ToGroupDto(&group)
}
