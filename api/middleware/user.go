package middleware

import (
	"app/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSKey = []byte("brandon02")

func GenerateJWT(user model.User, expiry time.Duration) (string, error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Uid,
		"exp":     exp,
	})

	tokenString, err := token.SignedString(JwtSKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateUser(newAuthor model.User) error {

	db := DB

	var existingUser model.User
	if err := db.Table("users").Where("user_name=?", newAuthor.UserName).First(&existingUser).Error; err == nil {
		return errors.New("Author_with_the_same_name_is_already_exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*model.User, error) {

	var user model.User
	if err := DB.Table("users").Where("user_email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
