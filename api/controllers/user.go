package controller

import (
	"app/middleware"
	"app/model"
	helpers "app/utility"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {

	var user model.User

	if err := c.Bind(&user); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_user")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPass)

	if err := model.CreateUser(user); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_create_user")
	}

	token, err := middleware.GenerateJWT(user, 24)
	if err != nil {
		fmt.Println(err, "<<<<<")
		return err
	}

	return c.JSON(http.StatusAccepted, map[string]any{
		"Token": token,
	})

}

func Login(c echo.Context) error {
	var userLogin model.User

	if err := c.Bind(&userLogin); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_user")
	}

	user, err := model.GetUserByUsername(userLogin.UserName)
	if err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_get_user_by_username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return helpers.ReturnLog(c, http.StatusUnauthorized, "Error_compare_password")
	}

	token, err := middleware.GenerateJWT(*user, 24*time.Hour)
	if err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_generate_token")
	}

	expTime := time.Now().Add(24 * time.Hour)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Token":  token,
		"Exp_at": expTime.Format(time.RFC3339),
	})
}
