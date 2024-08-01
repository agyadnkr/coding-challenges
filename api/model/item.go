package model

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

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
