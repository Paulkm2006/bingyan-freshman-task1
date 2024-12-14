package model

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/dto"
	"bingyan-freshman-task0/internal/service"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePost(post *dto.Post) (*dto.Post, error) {
	tx := db.Begin()

	content := post.Content
	post.Content = nil

	if err := tx.Clauses(clause.Returning{}).Create(post).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if content != nil {
		body := &dto.Body{
			PID:     post.PID,
			Content: *content,
		}
		if err := tx.Create(body).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := IncrArticle(post.NID); err != nil {
		tx.Rollback()
		return nil, err
	}

	err := service.IndexPost(post, content)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return post, nil
}

func GetPosts(paging param.Paging) ([]dto.Post, error) {
	var posts []dto.Post
	result := db.Find(&posts).
		Limit(paging.PageSize).
		Offset((paging.Page - 1) * paging.PageSize).
		Order(paging.SortingStatement())
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func GetPostByPID(pid int) (*dto.Post, error) {
	var post dto.Post
	result := db.Where("p_id = ?", pid).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	var body dto.Body
	res := db.Where("p_id = ?", pid).First(&body)
	if res.Error == nil {
		post.Content = &body.Content
	} else {
		return nil, res.Error
	}

	return &post, nil
}

func GetPostsByNID(paging param.Paging) ([]dto.Post, error) {
	var posts []dto.Post
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

func GetPostsByUID(paging param.Paging) ([]dto.Post, error) {
	var posts []dto.Post
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
	result := db.Where("p_id = ?", pid).Delete(&dto.Post{})
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("p_id = ?", pid).Delete(&dto.Like{})
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("p_id = ?", pid).Delete(&dto.Comment{})
	if result.Error != nil {
		return result.Error
	}
	result = db.Where("p_id = ?", pid).Delete(&dto.Body{})
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
	result := db.Model(&dto.Post{}).Where("p_id = ?", pid).Update("likes", gorm.Expr("likes + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DecrLikes(pid int) error {
	result := db.Model(&dto.Post{}).Where("p_id = ?", pid).Update("likes", gorm.Expr("likes - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IncrComments(pid int) error {
	result := db.Model(&dto.Post{}).Where("p_id = ?", pid).Update("comments", gorm.Expr("comments + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DecrComments(pid int) error {
	result := db.Model(&dto.Post{}).Where("p_id = ?", pid).Update("comments", gorm.Expr("comments - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Add this function to update post content
func UpdatePostContent(pid int, content string) error {
	body := &dto.Body{
		PID:     pid,
		Content: content,
	}
	result := db.Save(body)
	return result.Error
}
