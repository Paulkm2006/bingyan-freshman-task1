package model

import (
	"errors"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;index" query:"id"`
	Username string `json:"username" gorm:"unique" query:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Nickname string `json:"nickname,omitempty"`
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExist = errors.New("user already exist")

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
