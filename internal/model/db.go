package model

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := config.Config.DB.Dsn
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Post{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Comment{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Like{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Node{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Follow{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&dto.Body{}); err != nil {
		panic(err)
	}
}
