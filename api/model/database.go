package model

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var ErrDuplicateData = errors.New("duplicate_data")

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=challange port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
