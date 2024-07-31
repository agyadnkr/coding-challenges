package model

import (
	"app/model"
	"errors"

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

func FetchItem(request Filter) (interface{}, error) {

	var item Item
	db := DB

	queryBuilder := db.Table("items").Where("1=1")

	ItemsDetails := *model.Item

	if request.Filter != "" {
		filter := request.Filter
		queryBuilder = queryBuilder.Order(filter)

		return item, nil
	}

	if err := queryBuilder.Find(&item).Error; err != nil {
		return err, nil
	}

	return item, nil
}
