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
	product1 = entity.Product{
		ID:     "product1-f12e-4fa0-bf4f-ba002c11a671",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 1",
		Price:  200000000123.1234,
	}

	product2 = entity.Product{
		ID:     "product2-f12e-4fa0-bf4f-ba002c11a671",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 222222",
		Price:  1234.1234,
	}
	product3 = entity.Product{
		ID:     "product3-f12e-4fa0-bf4f-ba002c11a671",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 3",
		Price:  2000,
	}
)

func Test_productPg_CreateProduct(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewProductPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.CreateProduct(ctx, tx, product1)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}
