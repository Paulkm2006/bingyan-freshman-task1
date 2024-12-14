package model

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/dto"
)

func CreateComment(comment *dto.Comment) error {
	err := IncrComments(comment.PID)
	if err != nil {
		return err
	}
	result := db.Model(&dto.Comment{}).Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCommentByCID(cid int) (*dto.Comment, error) {
	var comment dto.Comment
	result := db.Where("c_id = ?", cid).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func GetCommentsByPID(paging param.Paging) ([]dto.Comment, error) {
	var comments []dto.Comment
	result := db.Where("p_id = ?", paging.Id).
		Find(&comments).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order("created desc")
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func GetCommentsByUID(paging param.Paging) ([]dto.Comment, error) {
	var comments []dto.Comment
	result := db.Where("uid = ?", paging.Id).
		Find(&comments).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order("created desc")
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func DeleteComment(cid int, pid int) error {
	err := DecrComments(pid)
	if err != nil {
		return err
	}
	result := db.Where("c_id = ?", cid).Delete(&dto.Comment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
