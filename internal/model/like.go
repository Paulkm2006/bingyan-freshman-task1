package model

import (
	"errors"
	"time"
)

type Like struct {
	LID     int       `json:"lid" gorm:"primaryKey;autoIncrement;index" query:"lid"`
	UID     int       `json:"uid" gorm:"index" query:"uid"`
	PID     int       `json:"pid" gorm:"index" query:"pid"`
	Created time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
}

var ErrLikeNotFound = errors.New("like not found")

func CreateLike(like *Like) error {
	err := IncrLikes(like.PID)
	if err != nil {
		return err
	}
	result := db.Model(&Like{}).Create(like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func LikeExist(uid, pid int) bool {
	var like Like
	result := db.Where("uid = ? AND p_id = ?", uid, pid).First(&like)
	return result.Error == nil
}

func GetLikesByUID(uid int) ([]Like, error) {
	var likes []Like
	result := db.Where("uid = ?", uid).Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}

func GetLikesByPID(pid int) ([]Like, error) {
	var likes []Like
	result := db.Where("p_id = ?", pid).Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}

func DeleteLike(uid, pid int) error {
	exist := LikeExist(uid, pid)
	if !exist {
		return ErrLikeNotFound
	}
	err := DecrLikes(pid)
	if err != nil {
		return err
	}
	result := db.Where("uid = ? AND p_id = ?", uid, pid).Delete(&Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
