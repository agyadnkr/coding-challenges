package model

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	Invid     string         `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Itmid     uint           `gorm:"column:item_id" json:"item_id"`
	Wid       string         `gorm:"column:warehouse_id" json:"warehouse_id"`
	Quantity  int            `gorm:"column:quantity" json:"quantity"`
}
