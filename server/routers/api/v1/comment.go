package v1

import (
	"github.com/gin-gonic/gin"
)

func Comments(c *gin.Context) {

	appG := app.Gin{C: c}

	articleID := com.StrTo(c.Param("article_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(articleID, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{ArticleID:articleID}
	comments, err := commentService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, comments)
}

func AddComment(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": make(map[string]interface{}),
	//})
}

func DelComment(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": make(map[string]string),
	//})
}
