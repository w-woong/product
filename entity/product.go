package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	NilProduct      = Product{}
	NilGroup        = Group{}
	NilGroupProduct = GroupProduct{}
)

type Product struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"<-"`
	ImgUrl      string     `json:"img_url" gorm:"type:string;size:2048"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256;not null"`
	Price       float64    `json:"price" gorm:"column:price;type:float;precision:16;scale:4;not null"`
	Description string     `json:"description" gorm:"column:description;type:string"`

	Groups GroupList `json:"groups,omitempty" gorm:"many2many:group_products"`
}

func (e *Product) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Product) CreateID() string {
	return uuid.New().String()
}

func (e *Product) CreateSetID() {
	e.ID = e.CreateID()
}

func (e Product) IsNil() bool {
	return e.ID == ""
}

type ProductList []Product

func (e *ProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Group struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"<-"`
	Name        string     `json:"name" gorm:"column:name;type:string;size:256;not null"`
	Description string     `json:"description" gorm:"column:description;type:string"`

	Products ProductList `json:"products,omitempty" gorm:"many2many:group_products;foreignKey:ID;joinForeignKey:GroupID;references:ID;joinReferences:ProductID"`
}

func (e *Group) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Group) CreateID() string {
	return uuid.New().String()
}

func (e *Group) CreateSetID() {
	e.ID = e.CreateID()
}

func (e Group) IsNil() bool {
	return e.ID == ""
}

type GroupList []Group

func (e *GroupList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type GroupProduct struct {
	GroupID   string     `json:"group_id" gorm:"primaryKey;type:string;size:64"`
	ProductID string     `json:"product_id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"<-:update"`
}

func (e *GroupProduct) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type GroupProductList []GroupProduct

func (e *GroupProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
