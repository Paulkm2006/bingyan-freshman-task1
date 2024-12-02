package model

import "time"

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

func GetCommentsByPID(pid int, page int, pageSize int) ([]Comment, error) {
	var comments []Comment
	result := db.Where("p_id = ?", pid).Find(&comments).Limit(pageSize).Offset((page - 1) * pageSize)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func GetCommentsByUID(uid int, page int, pageSize int) ([]Comment, error) {
	var comments []Comment
	result := db.Where("uid = ?", uid).Find(&comments).Limit(pageSize).Offset((page - 1) * pageSize)
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
