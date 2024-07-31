package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Uid       string         `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	UserName  string         `gorm:"column:username" json:"username"`
	Email     string         `gorm:"column:email" json:"email"`
	Password  string         `gorm:"column:password" json:"password"`
}

type Item struct {
	Itmid           string         `gorm:"column:id"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
	ItemName        string         `gorm:"column:name" json:"name"`
	ItemPrice       int            `gorm:"column:price" json:"price"`
	ItemDescription string         `gorm:"column:description" json:"description"`
}

type Filter struct {
	Filter   map[string]interface{} `json:"filter"`
	KeyWord  string                 `json:"keyword"`
	PriceMin float64                `json:"price_min"`
	PriceMax float64                `json:"price_max"`
	Itmid    string                 `json:"item_id"`
	Wid      string                 `json:"warehouse_id"`
}

type Warehouse struct {
	Wid              string         `gorm:"column:id"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
	WarehouseName    string         `gorm:"column:name" json:"name"`
	WarehouseAddress string         `gorm:"column:address" json:"address"`
}

type Inventory struct {
	Invid     string         `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Itmid     uint           `gorm:"column:item_id" json:"item_id"`
	Wid       string         `gorm:"column:warehouse_id" json:"warehouse_id"`
	Quantity  int            `gorm:"column:quantity" json:"quantity"`
}

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
}
