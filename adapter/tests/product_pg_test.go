package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/product/adapter"
	"github.com/w-woong/product/entity"
)

var (
	product1 = entity.Product{
		ID:     "test-product-1",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 1",
		Price:  200000000123.1234,
		Groups: entity.GroupList{
			group1,
		},
	}

	product2 = entity.Product{
		ID:     "test-product-2",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 222222",
		Price:  1234.1234,
		// Groups: entity.GroupList{
		// 	group1,
		// },
	}
	product3 = entity.Product{
		ID:     "test-product-3",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 3",
		Price:  2000,
		Groups: entity.GroupList{
			group1,
		},
	}
	product4 = entity.Product{
		ID:     "test-product-4",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 4",
		Price:  2000,
		Groups: entity.GroupList{
			group1,
		},
	}
	product5 = entity.Product{
		ID:     "test-product-5",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 5",
		Price:  2000,
		Groups: entity.GroupList{
			group1,
		},
	}
	product6 = entity.Product{
		ID:     "test-product-6",
		ImgUrl: "https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80",
		Name:   "product 6",
		Price:  2000,
		Groups: entity.GroupList{
			group1,
		},
	}
)

func Test_productPg_CreateProduct(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewProductPg(gdb)
	repoGroupProducts := adapter.NewGroupProductPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	repoGroupProducts.DeleteByProductID(ctx, tx, product1.ID)
	repoGroupProducts.DeleteByProductID(ctx, tx, product2.ID)
	repoGroupProducts.DeleteByProductID(ctx, tx, product3.ID)
	repoGroupProducts.DeleteByProductID(ctx, tx, product4.ID)
	repoGroupProducts.DeleteByProductID(ctx, tx, product5.ID)
	repoGroupProducts.DeleteByProductID(ctx, tx, product6.ID)

	repo.DeleteProduct(ctx, tx, product1.ID)
	repo.DeleteProduct(ctx, tx, product2.ID)
	repo.DeleteProduct(ctx, tx, product3.ID)
	repo.DeleteProduct(ctx, tx, product4.ID)
	repo.DeleteProduct(ctx, tx, product5.ID)
	repo.DeleteProduct(ctx, tx, product6.ID)

	affected, err := repo.CreateProduct(ctx, tx, product1)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	affected, err = repo.CreateProduct(ctx, tx, product2)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	affected, err = repo.CreateProduct(ctx, tx, product3)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	affected, err = repo.CreateProduct(ctx, tx, product4)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	affected, err = repo.CreateProduct(ctx, tx, product5)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	affected, err = repo.CreateProduct(ctx, tx, product6)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)
	assert.Nil(t, tx.Commit())
}

func Test_productPg_ReadProductNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	repo := adapter.NewProductPg(gdb)

	res, err := repo.ReadProductNoTx(ctx, product1.ID)
	assert.Nil(t, err)
	// assert.EqualValues(t, 1, affected)
	fmt.Println(res.String())

}
