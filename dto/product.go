package dto

import (
	"encoding/json"
	"time"
)

var (
	NilProduct      = Product{}
	NilProductGroup = ProductGroup{}
)

type Product struct {
	ID          string     `json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ImgUrl      string     `json:"img_url"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
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
	ID          string     `json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ProductID   string     `json:"product_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
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
