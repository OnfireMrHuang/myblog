package util

import (
	"server/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetLimit(c *gin.Context) int {
	limit, _ := com.StrTo(c.Query("limit")).Int()
	if limit < 1 {
		limit = setting.AppSetting.PageSize
	}
	return limit
}

func GetPage(c *gin.Context) int {
	result := 0
	pageSize := GetLimit(c)
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
