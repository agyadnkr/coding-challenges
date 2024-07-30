package model

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Uid        string    `gorm:"column:id"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
	Deleted_at gorm.DeletedAt
	UserName   string `gorm:"column:username" json:"username"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
}

type Item struct {
	Itmid           string    `gorm:"column:id"`
	Created_at      time.Time `gorm:"column:created_at"`
	Updated_at      time.Time `gorm:"column:updated_at"`
	Deleted_at      gorm.DeletedAt
	ItemName        string                 `gorm:"column:name" json:"name"`
	ItemPrice       int                    `gorm:"column:price" json:"price"`
	ItemDescription string                 `gorm:"column:description" json:"description"`
	ItmFilters      map[string]interface{} `json:"filter"`
}

type Warehouse struct {
	Wid              string    `gorm:"column:id"`
	Created_at       time.Time `gorm:"column:created_at"`
	Updated_at       time.Time `gorm:"column:updated_at"`
	Deleted_at       gorm.DeletedAt
	WarehouseName    string                 `gorm:"column:name" json:"name"`
	WarehouseAddress string                 `gorm:"column:address" json:"address"`
	WhFilters        map[string]interface{} `json:"filter"`
}

type Inventory struct {
	Invid      string    `gorm:"column:id"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
	Deleted_at gorm.DeletedAt
	Itmid      uint                   `gorm:"column:item_id" json:"item_id"`
	Wid        string                 `gorm:"column:warehouse_id" json:"warehouse_id"`
	Quantity   int                    `gorm:"column:quantity" json:"quantity"`
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
	if err := db.Table("users").Where("username=?", newAuthor.UserName).First(&existingUser).Error; err == nil {
		return errors.New("Author_with_the_same_name_is_already_exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByPassword(password string) (*User, error) {

	var user User
	if err := DB.Table("users").Where("password=?", password).First(&user).Error; err != nil {
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
