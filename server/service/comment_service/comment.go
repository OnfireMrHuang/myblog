package comment_service

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"

	"server/models"
	"server/pkg/export"
	"server/pkg/file"
	"server/pkg/gredis"
	"server/pkg/logging"
	"server/service/cache_service"
)

type Comment struct {
	ID         int
	Username string 
	Email    string 
	Content  string 
	IP       string 
	ArticleID int 

	PageNum  int
	PageSize int
}

// 添加评论
func (c *Comment) Add() error {
	comment := map[string]interface{}{
		"article_id":      c.ArticleID,
		"username":        c.Username,
		"email":           c.Email,
		"content":         c.Content,
		"ip":              c.IP,
	}
	ok := models.AddComment(comment)
	if !ok {
		return errors.New("add comment fail")
	}
	return nil
}

func (c *Comment) Delete() error {
	return models.DeleteComment(c.ID)
}

func (c *Comment) GetAll() ([]models.Comment, error) {
	var (
		comments []models.Comment
	)
	comments= models.GetComments(c.PageNum, c.PageSize, t.getMaps())
	return comments, nil
}

func (c *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if c.ArticleID > 0 {
		maps["article_id"] = c.ArticleID
	}
	return maps
}