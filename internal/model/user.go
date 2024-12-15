package model

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"
	"crypto/md5"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExist = errors.New("user already exist")

func AddDefaultAdmin() error {
	// Add default admin
	_, err := GetUserByUsername("admin")
	if err == nil {
		return ErrUserAlreadyExist
	}
	admin := &dto.User{
		Username:   config.Config.Admin.Username,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(config.Config.Admin.Password))),
		Permission: 1,
	}
	resultUser := db.Model(&dto.User{}).Create(admin)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	return nil
}

func AddUser(user *dto.User) error {
	// Add user
	_, err := GetUserByUsername(user.Username)
	if err == nil {
		return ErrUserAlreadyExist
	}
	resultUser := db.Model(&dto.User{}).Create(user)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	return nil

}

func UpdateUser(user *dto.User) error {
	// Update user
	var old dto.User
	result := db.Where("id = ?", user.ID).First(&old)
	if result.Error != nil {
		return result.Error
	}
	user.ID = old.ID
	result = db.Save(user)
	return result.Error
}

func DeleteUser(id int) error {
	// Delete user
	_, err := GetUserByID(id)
	if err != nil {
		return err
	}
	result := db.Delete(&dto.User{}, id)
	return result.Error
}

func GetUsers() ([]dto.User, error) {
	// Get users
	users := []dto.User{}
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(id int) (*dto.User, error) {
	// Get user
	user := &dto.User{}
	result := db.Where("id = ?", id).First(user)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, ErrUserNotFound
	}
	return user, result.Error
}

func GetUserByUsername(username string) (*dto.User, error) {
	// Get user
	user := &dto.User{}
	result := db.Where("username = ?", username).First(user)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, ErrUserNotFound
	}
	return user, result.Error
}

func IncrFollowed(id int) error {
	// Increase followed
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	user.Followed++
	result := db.Save(user)
	return result.Error
}

func DecrFollowed(id int) error {
	// Decrease followed
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	user.Followed--
	result := db.Save(user)
	return result.Error
}

func IncrFollowers(id int) error {
	// Increase followers
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	user.Followers++
	result := db.Save(user)
	return result.Error
}

func DecrFollowers(id int) error {
	// Decrease followers
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	user.Followers--
	result := db.Save(user)
	return result.Error
}
