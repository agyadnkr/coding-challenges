package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

type CreateInventoryRequest struct {
	WarehouseID string          `json:"warehouse_id"`
	Items       []InventoryItem `json:"items"`
}

type InventoryItem struct {
	ItemID   string  `json:"item_id"`
	Quantity float64 `json:"quantity"`
}

type UpdateInventoryRequest struct {
	WarehouseID string          `json:"warehouse_id"`
	Items       []InventoryItem `json:"items"`
}

type StockMoveRequest struct {
	OriginWid      string          `json:"origin_warehouse_id"`
	DestinationWid string          `json:"destination_warehouse_id"`
	Items          []StockMoveItem `json:"items"`
}

type StockMoveItem struct {
	ItemID   string  `json:"item_id"`
	Quantity float64 `json:"quantity"`
}

var ErrDuplicatedData = errors.New("duplicated_data")
var ErrItemNotFound = errors.New("item_not_found")
var ErrNotEnoughStock = errors.New("not_enough_stock_in_origin_warehouse")

func CreateInventory(req CreateInventoryRequest) error {
	tx := DB.Begin()

	for _, item := range req.Items {
		var existingInventory Inventory
		if err := tx.Where("warehouse_id = ? AND item_id = ?", req.WarehouseID, item.ItemID).First(&existingInventory).Error; err == nil {
			tx.Rollback()
			return ErrDuplicatedData
		}

		newInventory := Inventory{
			Invid:    uuid.New().String(),
			Wid:      req.WarehouseID,
			Itmid:    item.ItemID,
			Quantity: int(item.Quantity),
		}

		if err := tx.Create(&newInventory).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func FetchInventories(request Filter) ([]Inventory, error) {
	var inventories []Inventory
	db := DB

	queryBuilder := db.Where("deleted_at IS NULL")

	if request.Filter != nil {
		if itemID, ok := request.Filter["item_id"]; ok && itemID != "" {
			queryBuilder = queryBuilder.Where("item_id = ?", itemID)
		}
		if warehouseID, ok := request.Filter["warehouse_id"]; ok && warehouseID != "" {
			queryBuilder = queryBuilder.Where("warehouse_id = ?", warehouseID)
		}
	}

	if err := queryBuilder.Find(&inventories).Error; err != nil {
		return nil, err
	}

	return inventories, nil
}

func UpdateInventory(inventoryID string, req UpdateInventoryRequest) error {
	tx := DB.Begin()

	for _, item := range req.Items {
		var existingInventory Inventory
		if err := tx.Where("warehouse_id = ? AND item_id = ?", req.WarehouseID, item.ItemID).First(&existingInventory).Error; err != nil {
			tx.Rollback()
			return ErrItemNotFound
		}

		newQuantity := float64(existingInventory.Quantity) + item.Quantity
		if err := tx.Model(&Inventory{}).Where("id = ?", existingInventory.Invid).Update("quantity", newQuantity).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func MoveStock(stockMoveRequest StockMoveRequest) error {
	tx := DB.Begin()

	for _, item := range stockMoveRequest.Items {
		var originInventory, destinationInventory Inventory

		if err := tx.Model(&Inventory{}).Where("warehouse_id = ? AND item_id = ?", stockMoveRequest.OriginWid, item.ItemID).First(&originInventory).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("inventory_record_not_found")
			}
			return err
		}

		if originInventory.Quantity < int(item.Quantity) {
			tx.Rollback()
			return errors.New("not_enough_stock_in_origin_warehouse")
		}

		if err := tx.Model(&Inventory{}).Where("id = ?", originInventory.Invid).Update("quantity", gorm.Expr("quantity - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&Inventory{}).Where("warehouse_id = ? AND item_id = ?", stockMoveRequest.DestinationWid, item.ItemID).First(&destinationInventory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newInventory := Inventory{
					Invid:    uuid.New().String(),
					Itmid:    item.ItemID,
					Wid:      stockMoveRequest.DestinationWid,
					Quantity: int(item.Quantity),
				}

				if err := tx.Create(&newInventory).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else {
				tx.Rollback()
				return err
			}
		} else {
			if err := tx.Model(&Inventory{}).Where("id = ?", destinationInventory.Invid).Update("quantity", gorm.Expr("quantity + ?", item.Quantity)).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
