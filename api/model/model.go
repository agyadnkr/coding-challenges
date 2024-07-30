package model

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Uid        string `gorm:"user_id" column:"user_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	UserName   string `gorm:"username" column:"username"`
	Email      string `gorm:"email" column:"email"`
	Password   string `gorm:"password" column:"password"`
}

type Item struct {
	Itmid           string `gorm:"item_id" column:"item_id"`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      gorm.DeletedAt
	ItemName        string                 `gorm:"name" column:"name"`
	ItemPrice       int                    `gorm:"price" column:"price"`
	ItemDescription string                 `gorm:"description" column:"description"`
	ItmFilters      map[string]interface{} `json:"filter"`
}

type Warehouse struct {
	Wid              string `gorm:"warehouse_id" column:"warehouse_id"`
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       gorm.DeletedAt
	WarehouseName    string                 `gorm:"name" column:"name"`
	WarehouseAddress string                 `gorm:"address" column:"address"`
	WhFilters        map[string]interface{} `json:"filter"`
}

type Inventory struct {
	Invid      string `gorm:"inventory_id" column:"inventory_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Itmid      uint                   `gorm:"item_id" column:"item_id"`
	Wid        string                 `gorm:"warehouse_id" column:"warehouse_id"`
	Quantity   int                    `gorm:"quantity" column:"quantity"`
	InvFilters map[string]interface{} `json:"filter"`
}

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=challange port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateUser(newAuthor User) error {

	db := DB

	var existingUser User
	if err := db.Table("users").Where("name=?", newAuthor.UserName).First(&existingUser).Error; err == nil {
		return errors.New("Author_with_the_same_name_is_already_exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {

	var user User
	if err := DB.Table("users").Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
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

	return nil

}
