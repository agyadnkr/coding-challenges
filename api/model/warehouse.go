package model

import (
	"errors"
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

func CreateWarehouse(newWarehouse Warehouse) error {
	db := DB

	var existingWarehouse Warehouse
	if err := db.Table("warehouses").Where("name = ?", newWarehouse.WarehouseName).First(&existingWarehouse).Error; err == nil {
		return errors.New("Warehouse_with_the_same_name_already_exists")
	}

	if err := db.Create(&newWarehouse).Error; err != nil {
		return err
	}

	return nil
}
