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
	NID      int       `json:"nid" gorm:"index" query:"nid"`
	Likes    int       `json:"likes" gorm:"default:0"`
	Comments int       `json:"comments" gorm:"default:0"`
	Content  *string   `json:"content,omitempty" gorm:"-"` // Exclude from database
}

type Body struct {
	PID     int    `json:"pid" gorm:"primaryKey;index" query:"pid"`
	Content string `json:"content" query:"content"`
}

func CreatePost(post *Post) (*Post, error) {
	tx := db.Begin()

	content := post.Content
	post.Content = nil

	if err := tx.Clauses(clause.Returning{}).Create(post).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if content != nil {
		body := &Body{
			PID:     post.PID,
			Content: *content,
		}
		if err := tx.Create(body).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := IncrArticle(post.PID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return post, nil
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
	if result.Error != nil {
		return nil, result.Error
	}
	var body Body
	res := db.Where("p_id = ?", pid).First(&body)
	if res.Error == nil {
		post.Content = &body.Content
	} else {
		return nil, res.Error
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
	result = db.Where("p_id = ?", pid).Delete(&Body{})
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

// Add this function to update post content
func UpdatePostContent(pid int, content string) error {
	body := &Body{
		PID:     pid,
		Content: content,
	}
	result := db.Save(body)
	return result.Error
}
