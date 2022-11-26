package dto

import (
	"encoding/json"
	"time"
)

var (
	NilProduct = Product{}
	NilGroup   = Group{}
)

type Product struct {
	ID          string     `json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ImgUrl      string     `json:"img_url"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`

	Groups GroupList `json:"groups,omitempty"`
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

type Group struct {
	ID          string     `json:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`

	Products ProductList `json:"products,omitempty"`
}

func (e *Group) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type GroupList []Group

func (e *GroupList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
