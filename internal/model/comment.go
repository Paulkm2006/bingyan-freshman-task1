package model

import (
	"bingyan-freshman-task0/internal/controller/param"
	"time"
)

type Comment struct {
	CID     int       `json:"cid" gorm:"primaryKey;autoIncrement;index" query:"cid"`
	UID     int       `json:"uid" gorm:"index" query:"uid"`
	Content string    `json:"content" query:"content"`
	PID     int       `json:"pid" gorm:"index" query:"pid"`
	Created time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
}

func CreateComment(comment *Comment) error {
	err := IncrComments(comment.PID)
	if err != nil {
		return err
	}
	result := db.Model(&Comment{}).Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCommentByCID(cid int) (*Comment, error) {
	var comment Comment
	result := db.Where("c_id = ?", cid).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func GetCommentsByPID(paging param.Paging) ([]Comment, error) {
	var comments []Comment
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

func GetCommentsByUID(paging param.Paging) ([]Comment, error) {
	var comments []Comment
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
	result := db.Where("c_id = ?", cid).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
