package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Uid       string         `gorm:"column:id"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	UserName  string         `gorm:"column:username" json:"username"`
	Email     string         `gorm:"column:email" json:"email"`
	Password  string         `gorm:"column:password" json:"password"`
}

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

var ErrDuplicateData = errors.New("duplicate_data")

func CreateUser(newUser User) error {
	db := DB

	var existingUser User
	if err := db.Table("users").Where("username = ? OR email = ?", newUser.UserName, newUser.Email).First(&existingUser).Error; err == nil {
		return ErrDuplicateData
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	newUser.Uid = uuid.New().String()

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := DB.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
