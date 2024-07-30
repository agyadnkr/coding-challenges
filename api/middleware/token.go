package middleware

import (
	"app/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var jwtSecret string

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not find .env file", err)
	}

	jwtSecret = viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in .env file")
	}
}

func NewEnv() *model.Env {
	env := model.Env{}
	viper.SetConfigFile("key.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not find Env file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Enviroment cant be loaded", err)
	}

	if env.AppEnv == "development" {
		log.Println("The app is running in development env")
	}

	return &env

}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "No_token_provided",
			})
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid_token_format",
			})
		}

		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid_token",
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid_token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid_token",
			})
		}

		c.Set("user_email", claims["user_email"])
		c.Set("user_name", claims["user_name"])

		return next(c)
	}
}

func GenerateJWT(user model.User, expiry time.Duration) (string, error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Uid,
		"exp":     exp,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
