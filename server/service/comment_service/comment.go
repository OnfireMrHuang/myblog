package comment_service

import (
	"errors"
	"server/models"
)

type Comment struct {
	ID        int
	Username  string
	Email     string
	Content   string
	IP        string
	ArticleID int

	PageNum  int
	PageSize int
}

// 添加评论
func (c *Comment) Add() error {
	comment := map[string]interface{}{
		"article_id": c.ArticleID,
		"username":   c.Username,
		"email":      c.Email,
		"content":    c.Content,
		"ip":         c.IP,
	}
	return models.AddComment(comment)
}

func (c *Comment) Delete() error {
	ok := models.DeleteComment(c.ID)
	if !ok {
		return errors.New("删除评论记录失败")
	}
	return nil
}

func (c *Comment) GetAll() ([]models.Comment, error) {
	var (
		comments []models.Comment
	)
	comments = models.GetComments(c.PageNum, c.PageSize, c.getMaps())
	return comments, nil
}

func (c *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if c.ArticleID > 0 {
		maps["article_id"] = c.ArticleID
	}
	return maps
}
