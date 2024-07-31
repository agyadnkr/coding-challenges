package model

import (
	"errors"

	"github.com/google/uuid"
)

func CreateUser(newAuthor User) error {

	db := DB

	var existingUser User
	if err := db.Table("users").Where("username=?", newAuthor.UserName).First(&existingUser).Error; err == nil {
		return errors.New("Author_with_the_same_name_is_already_exists")
	}

	newAuthor.Uid = uuid.New().String()

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
