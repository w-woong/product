package adapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/product/adapter"
	"github.com/w-woong/product/entity"
)

var (
	group1 = entity.Group{
		ID:          "group001-f12e-4fa0-bf4f-ba002c11a671",
		Name:        "Group 1",
		Description: "Group 1 Description",
	}
	group2 = entity.Group{
		ID:          "group002-f12e-4fa0-bf4f-ba002c11a671",
		Name:        "Group 2",
		Description: "Group 2 Description",
		Products: entity.ProductList{
			product2,
		},
	}
)

func Test_groupPg_CreateGroup(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewGroupPg(gdb)
	repoGP := adapter.NewGroupProductPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	repoGP.DeleteByGroupID(ctx, tx, group1.ID)
	repo.DeleteGroup(ctx, tx, group1.ID)

	affected, err := repo.CreateGroup(ctx, tx, group1)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

func Test_groupPg_CreateGroup2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewGroupPg(gdb)
	repoGP := adapter.NewGroupProductPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	repoGP.DeleteByGroupID(ctx, tx, group2.ID)
	repo.DeleteGroup(ctx, tx, group2.ID)

	affected, err := repo.CreateGroup(ctx, tx, group2)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

// func Test_groupPg_ReadGroup(t *testing.T) {
// 	if !onlinetest {
// 		t.Skip("skipping online tests")
// 	}

// 	ctx := context.Background()

// 	beginner := txcom.NewGormTxBeginner(gdb)
// 	repo := adapter.NewGroupPg(gdb)

// 	tx, err := beginner.Begin()
// 	assert.Nil(t, err)
// 	defer tx.Rollback()

// 	res, err := repo.ReadGroup(ctx, tx, group1.ID)
// 	assert.Nil(t, err)

// 	expected := `{"id":"group001-f12e-4fa0-bf4f-ba002c11a671","created_at":"2022-11-26T09:30:34.164683+09:00","name":"Group 1","description":"Group 1 Description","products":[{"id":"product1-f12e-4fa0-bf4f-ba002c11a671","created_at":"2022-11-26T09:30:34.163071+09:00","img_url":"https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5\u0026ixid=eyJhcHBfaWQiOjEyMDd9\u0026s=6ff92caffcdd63681a35134a6770ed3b\u0026auto=format\u0026fit=crop\u0026w=1951\u0026q=80","name":"product 1","price":200000000123.1234,"description":""}]}`
// 	assert.EqualValues(t, expected, res.String())

// 	assert.Nil(t, tx.Commit())

// 	res, err = repo.ReadGroupNoTx(ctx, group1.ID)
// 	assert.Nil(t, err)
// 	assert.EqualValues(t, expected, res.String())
// }

func Test_groupPg_UpdateGroup(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewGroupPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	group2.Description = "Updated description"
	affected, err := repo.UpdateGroup(ctx, tx, group2.ID, group2)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

func Test_groupPg_DeleteGroup(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewGroupPg(gdb)
	repoGP := adapter.NewGroupProductPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	repoGP.DeleteByGroupID(ctx, tx, group2.ID)

	affected, err := repo.DeleteGroup(ctx, tx, group2.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
