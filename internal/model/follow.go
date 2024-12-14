package model

import "bingyan-freshman-task0/internal/dto"

func CreateFollow(follow *dto.Follow) error {
	result := db.Model(&dto.Follow{}).Create(follow)
	if result.Error != nil {
		return result.Error
	}
	err := IncrFollowed(follow.UID)
	if err != nil {
		return err
	}
	err = IncrFollowers(follow.Followee)
	if err != nil {
		return err
	}
	return nil
}

func GetFollowsByUID(uid int) ([]int, error) {
	var follows []int
	result := db.Model(&dto.Follow{}).Where("uid = ?", uid).Pluck("followee", &follows)
	if result.Error != nil {
		return nil, result.Error
	}
	return follows, nil
}

func GetFollowersByUID(uid int) ([]int, error) {
	var followers []int
	result := db.Model(&dto.Follow{}).Where("followee = ?", uid).Pluck("uid", &followers)
	if result.Error != nil {
		return nil, result.Error
	}
	return followers, nil
}

func DeleteFollow(uid, followee int) error {
	result := db.Where("uid = ? AND followee = ?", uid, followee).Delete(&dto.Follow{})
	if result.Error != nil {
		return result.Error
	}
	err := DecrFollowed(uid)
	if err != nil {
		return err
	}
	err = DecrFollowers(followee)
	if err != nil {
		return err
	}
	return nil
}
