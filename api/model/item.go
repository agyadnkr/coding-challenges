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
	ItemPrice       int            `gorm:"column:price" json:"price"`
	ItemDescription string         `gorm:"column:description" json:"description"`
}

func CreateItem(newItem Item) error {

	db := DB

	var existingItem Item

	if err := db.Table("items").Where("name=?", newItem.ItemName).First(&existingItem).Error; err == nil {
		return errors.New("Item_with_the_same_name_already_exists")
	}

	if err := db.Create(&newItem).Error; err != nil {
		return err
	}

	existingItem.Itmid = uuid.New().String()

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

func UpdateItem(itemID string, UpdateItem Item) error {
	if err := DB.Model(&Item{}).Where("id = ?", itemID).Updates(UpdateItem).Error; err != nil {
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
