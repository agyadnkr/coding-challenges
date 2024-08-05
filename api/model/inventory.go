package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	Invid     string         `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Itmid     string         `gorm:"column:item_id" json:"item_id"`
	Wid       string         `gorm:"column:warehouse_id" json:"warehouse_id"`
	Quantity  int            `gorm:"column:quantity" json:"quantity"`
}

type StockMoveRequest struct {
	OriginWid      string `json:"origin_warehouse_id"`
	DestinationWid string `json:"destination_warehouse_id"`
	Items          []struct {
		Itmid    string `json:"item_id"`
		Quantity int    `json:"quantity"`
	} `json:"items"`
}

func CreateInventory(newInventory *Inventory) error {
	db := DB

	var existingInventory Inventory
	if err := db.Table("inventories").Where("item_id = ? AND warehouse_id = ?", newInventory.Itmid, newInventory.Wid).First(&existingInventory).Error; err == nil {
		return errors.New("Inventory_for_this_item_and_warehouse_already_exists")
	}

	if err := db.Create(&newInventory).Error; err != nil {
		return err
	}

	return nil
}

func GetAllInventories() ([]Inventory, error) {
	var inventories []Inventory
	if err := DB.Where("deleted_at IS NULL").Find(&inventories).Error; err != nil {
		return nil, err
	}

	return inventories, nil
}

func UpdateInventory(inventoryID string, updatedInventory Inventory) error {
	if err := DB.Model(&Inventory{}).Where("id = ?", inventoryID).Updates(updatedInventory).Error; err != nil {
		return err
	}

	return nil
}

func DeleteInventory(inventoryID string) error {
	if err := DB.Model(&Inventory{}).Where("id = ?", inventoryID).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func MoveStock(stockMoveRequest StockMoveRequest) error {

	return nil
}
