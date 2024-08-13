package model

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Itmid           string         `gorm:"column:id"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
	ItemName        string         `gorm:"column:name" json:"name"`
	ItemPrice       float64        `gorm:"column:price" json:"price"`
	ItemDescription string         `gorm:"column:description" json:"description"`
}

func CreateItem(itemID string, newItem *Item) error {
	db := DB

	var existingItem, item Item

	if err := db.Table("items").First(&item, "id = ?", itemID).Error; err != nil {
		return nil
	}

	if err := db.Table("items").Where("name = ?", newItem.ItemName).First(&existingItem).Error; err == nil {
		return nil
	}

	newItem.Itmid = uuid.New().String()
	if err := db.Create(newItem).Error; err != nil {
		return nil
	}

	return nil
}

func CreateItems(newItem *Item) error {
	db := DB

	var existingItem Item
	if err := db.Table("items").Where("name = ?", newItem.ItemName).First(&existingItem).Error; err == nil {
		return errors.New("item_with_the_same_name_already_exists")
	}

	newItem.Itmid = uuid.New().String()
	if err := db.Create(newItem).Error; err != nil {
		return err
	}

	return nil
}

func FetchItem(request Filter) ([]Item, error) {
	var items []Item
	db := DB

	queryBuilder := db.Table("items").Where("deleted_at IS NULL")

	if request.KeyWord != "" {
		queryBuilder = queryBuilder.Where("name ILIKE ?", "%"+strings.TrimSpace(request.KeyWord)+"%")
	}

	if request.PriceMin != 0 {
		queryBuilder = queryBuilder.Where("price >= ?", request.PriceMin)
	}

	if request.PriceMax != 0 {
		queryBuilder = queryBuilder.Where("price <= ?", request.PriceMax)
	}

	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func ViewItem(itemID string) (*Item, error) {
	db := DB

	var item Item
	if err := db.Table("items").First(&item, "id = ?", itemID).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateItem(itemID string, updatedItem Item) error {
	if err := DB.Model(&Item{}).Where("id = ?", itemID).Updates(map[string]interface{}{
		"name":        updatedItem.ItemName,
		"price":       updatedItem.ItemPrice,
		"description": updatedItem.ItemDescription,
		"updated_at":  updatedItem.UpdatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteItem(itemID string) error {
	if err := DB.Where("id = ?", itemID).Delete(&Item{}).Error; err != nil {
		return err
	}

	return nil
}
