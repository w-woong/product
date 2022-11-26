package conv

import (
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.Product{}, &entity.Product{})
	structmapper.StoreMapper(&entity.Product{}, &dto.Product{})

	structmapper.StoreMapper(&dto.Group{}, &entity.Group{})
	structmapper.StoreMapper(&entity.Group{}, &dto.Group{})
}

func ToProductEntity(src *dto.Product) (res entity.Product, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToProductDto(src *entity.Product) (res dto.Product, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToGroupEntity(src *dto.Group) (res entity.Group, err error) {
	err = structmapper.Map(src, &res)
	return
}

func ToGroupDto(src *entity.Group) (res dto.Group, err error) {
	err = structmapper.Map(src, &res)
	return
}
