package entity

import (
	"encoding/json"
	"time"
)

var (
	NilProduct = Product{}
)

type Product struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"<-:update"`
	ImgUrl      string     `json:"img_url" gorm:"type:string;size:2048"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256;not null"`
	Price       float64    `json:"price" gorm:"column:price;type:float;precision:16;scale:4;not null"`
	Description string     `json:"description" gorm:"column:description;type:string"`
}

func (e *Product) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ProductList []Product

func (e *ProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ProductGroup struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"<-:update"`
	ProductID   string     `json:"product_id" gorm:"not null;type:string;size:64"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256;not null"`
	Description string     `json:"description" gorm:"column:description;type:string"`
}

func (e *ProductGroup) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type ProductGroupList []ProductGroup

func (e *ProductGroupList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
