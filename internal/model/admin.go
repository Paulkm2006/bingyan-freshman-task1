package model

import (
	"bingyan-freshman-task0/internal/config"
	"crypto/md5"
	"fmt"
)

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func AddDefaultAdmin() {
	admin := &Admin{
		Username: config.Config.Admin.Username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(config.Config.Admin.Password))),
	}
	db.Model(&Admin{}).Create(admin)
}

func AddAdmin(adminUser *Admin) error {
	result := db.Model(&Admin{}).Create(adminUser)
	return result.Error
}

func UpdateAdmin(adminUser *Admin) error {
	result := db.First(&Admin{}, adminUser.ID)
	if result.Error != nil {
		return result.Error
	}
	result = db.Save(adminUser)
	return result.Error
}

func DeleteAdmin(id int) error {
	result := db.Delete(&Admin{}, id)
	return result.Error
}

func GetAdmin(username string) (*Admin, error) {
	adminUser := &Admin{}
	result := db.Where("username = ?", username).First(adminUser)
	return adminUser, result.Error
}
