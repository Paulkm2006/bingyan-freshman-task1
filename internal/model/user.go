package model

import (
	"bingyan-freshman-task0/internal/config"
	"crypto/md5"
	"errors"
	"fmt"
)

type User struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;index" query:"id"`
	Username   string `json:"username" gorm:"unique" query:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname,omitempty"`
	Permission int    `json:"permission" gorm:"default:0"`
	Followed   int    `json:"followed" gorm:"default:0"`  // Number of people followed
	Followers  int    `json:"followers" gorm:"default:0"` // Number of followers
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExist = errors.New("user already exist")

func AddDefaultAdmin() error {
	// Add default admin
	_, err := GetUserByUsername("admin")
	if err == nil {
		return ErrUserAlreadyExist
	}
	admin := &User{
		Username:   config.Config.Admin.Username,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(config.Config.Admin.Password))),
		Permission: 1,
	}
	resultUser := db.Model(&User{}).Create(admin)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	return nil
}

func AddUser(user *User) error {
	// Add user
	_, err := GetUserByUsername(user.Username)
	if err == nil {
		return ErrUserAlreadyExist
	}
	resultUser := db.Model(&User{}).Create(user)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	return nil

}

func UpdateUser(user *User) error {
	// Update user
	var old User
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
	result := db.Delete(&User{}, id)
	return result.Error
}

func GetUserByID(id int) (*User, error) {
	// Get user
	user := &User{}
	result := db.Where("id = ?", id).First(user)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, ErrUserNotFound
	}
	return user, result.Error
}

func GetUserByUsername(username string) (*User, error) {
	// Get user
	user := &User{}
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
