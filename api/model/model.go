package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Uid        uint `gorm:"user_id" column:"user_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	UserName   string `gorm:"user_name" column:"user_name"`
	Email      string `gorm:"user_email" column:"user_email"`
	Password   string `gorm:"user_password" column:"user_password"`
}

type Item struct {
	Itmid           uint `gorm:"item_id" column:"item_id"`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      gorm.DeletedAt
	ItemName        string `gorm:"item_name" column:"item_name"`
	ItemPrice       int    `gorm:"item_price" column:"item_price"`
	ItemDescription string `gorm:"item_description" column:"item_description"`
}

type Warehouse struct {
	Wid              uint `gorm:"warehouse_id" column:"warehouse_id"`
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       gorm.DeletedAt
	WarehouseName    string `gorm:"warehouse_name" column:"warehouse_name"`
	WarehouseAddress string `gorm:"warehouse_address" column:"warehouse_address"`
}

type Inventory struct {
	Invid      uint `gorm:"inventory_id" column:"inventory_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Itmid      uint `gorm:"item_id" column:"item_id"`
	Wid        uint `gorm:"warehouse_id" column:"warehouse_id"`
	Quantity   int  `gorm:"item_quantity" column:"item_quantity"`
}
