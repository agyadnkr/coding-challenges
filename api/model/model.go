package model

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Uid        uint `gorm:"user_id" column:"user_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	UserName   string `gorm:"user_name" column:"user_name"`
	Email      string `gorm:"user_email" column:"user_email"`
	Password   string `gorm:"user_password" column:"user_password"`
}

type Item struct {
	Itmid           uint `gorm:"item_id" column:"item_id"`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      gorm.DeletedAt
	ItemName        string `gorm:"item_name" column:"item_name"`
	ItemPrice       int    `gorm:"item_price" column:"item_price"`
	ItemDescription string `gorm:"item_description" column:"item_description"`
}

type Warehouse struct {
	Wid              uint `gorm:"warehouse_id" column:"warehouse_id"`
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       gorm.DeletedAt
	WarehouseName    string `gorm:"warehouse_name" column:"warehouse_name"`
	WarehouseAddress string `gorm:"warehouse_address" column:"warehouse_address"`
}

type Inventory struct {
	Invid      uint `gorm:"inventory_id" column:"inventory_id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Itmid      uint `gorm:"item_id" column:"item_id"`
	Wid        uint `gorm:"warehouse_id" column:"warehouse_id"`
	Quantity   int  `gorm:"item_quantity" column:"item_quantity"`
}

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateUser(newAuthor User) error {

	db := DB

	var existingUser User
	if err := db.Table("users").Where("user_name=?", newAuthor.UserName).First(&existingUser).Error; err == nil {
		return errors.New("Author_with_the_same_name_is_already_exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {

	var user User
	if err := DB.Table("users").Where("user_email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
