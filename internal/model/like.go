package model

import (
	"bingyan-freshman-task0/internal/dto"
	"errors"
)

var ErrLikeNotFound = errors.New("like not found")

func CreateLike(like *dto.Like) error {
	err := IncrLikes(like.PID)
	if err != nil {
		return err
	}
	result := db.Model(&dto.Like{}).Create(like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func LikeExist(uid, pid int) bool {
	var like dto.Like
	result := db.Where("uid = ? AND p_id = ?", uid, pid).First(&like)
	return result.Error == nil
}

func GetLikesByUID(uid int) ([]dto.Like, error) {
	var likes []dto.Like
	result := db.Where("uid = ?", uid).Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}

func GetLikesByPID(pid int) ([]dto.Like, error) {
	var likes []dto.Like
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
	result := db.Where("uid = ? AND p_id = ?", uid, pid).Delete(&dto.Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
