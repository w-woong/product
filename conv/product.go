package conv

import (
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.Product{}, &entity.Product{})
	structmapper.StoreMapper(&entity.Product{}, &dto.Product{})

	structmapper.StoreMapper(&dto.ProductGroup{}, &entity.ProductGroup{})
	structmapper.StoreMapper(&entity.ProductGroup{}, &dto.ProductGroup{})
}

func ToProductEntity(src *dto.Product) (res entity.Product, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToProductDto(src *entity.Product) (res dto.Product, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToProductGroupEntity(src *dto.ProductGroup) (res entity.ProductGroup, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToProductGroupDto(src *entity.ProductGroup) (res dto.ProductGroup, err error) {
	err = structmapper.Map(src, &res)
	return
}
