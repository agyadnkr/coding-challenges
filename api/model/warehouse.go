package model

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	Wid              string         `gorm:"column:id"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
	WarehouseName    string         `gorm:"column:name" json:"name"`
	WarehouseAddress string         `gorm:"column:address" json:"address"`
}
