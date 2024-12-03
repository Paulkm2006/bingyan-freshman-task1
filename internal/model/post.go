package model

import (
	"bingyan-freshman-task0/internal/controller/param"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Post struct {
	PID      int       `json:"pid" gorm:"primaryKey;autoIncrement;index" query:"pid"`
	Created  time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
	UID      int       `json:"uid" gorm:"index" query:"uid"`
	Title    string    `json:"title" query:"title"`
	Content  string    `json:"content" query:"content"`
	NID      int       `json:"nid" gorm:"index" query:"nid"`
	Likes    int       `json:"likes" gorm:"default:0"`
	Comments int       `json:"comments" gorm:"default:0"`
}

func CreatePost(post *Post) (int, error) {
	result := db.Model(&Post{}).Clauses(clause.Returning{}).Create(post)
	if result.Error != nil {
		return 0, result.Error
	}
	err := IncrArticle(post.NID)
	if err != nil {
		return 0, err
	}
	return post.PID, nil
}

func GetPosts(paging param.Paging) ([]Post, error) {
	var posts []Post
	result := db.Find(&posts).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order(paging.SortingStatement())
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func GetPostByPID(pid int) (*Post, error) {
	var post Post
	result := db.Where("p_id = ?", pid).First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &post, nil
}

func GetPostsByNID(paging param.Paging) ([]Post, error) {
	var posts []Post
	result := db.Where("n_id = ?", paging.Id).
		Find(&posts).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order(paging.SortingStatement())
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func GetPostsByUID(paging param.Paging) ([]Post, error) {
	var posts []Post
	result := db.Where("uid = ?", paging.Id).
		Find(&posts).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order(paging.SortingStatement())
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func DeletePost(pid int) error {
	result := db.Where("p_id = ?", pid).Delete(&Post{})
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("p_id = ?", pid).Delete(&Like{})
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("p_id = ?", pid).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}
	err := DecrArticle(pid)
	if err != nil {
		return err
	}
	return nil
}

func IncrLikes(pid int) error {
	result := db.Model(&Post{}).Where("p_id = ?", pid).Update("likes", gorm.Expr("likes + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DecrLikes(pid int) error {
	result := db.Model(&Post{}).Where("p_id = ?", pid).Update("likes", gorm.Expr("likes - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IncrComments(pid int) error {
	result := db.Model(&Post{}).Where("p_id = ?", pid).Update("comments", gorm.Expr("comments + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DecrComments(pid int) error {
	result := db.Model(&Post{}).Where("p_id = ?", pid).Update("comments", gorm.Expr("comments - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
