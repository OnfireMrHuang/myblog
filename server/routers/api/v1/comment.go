package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"server/pkg/app"
	"server/pkg/e"
	"server/service/comment_service"
)

func Comments(c *gin.Context) {
	appG := app.Gin{C: c}
	var article struct {
		ArticleID int `json:"article_id"`
	}
	err := c.BindJSON(&article)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	articleID := article.ArticleID
	valid := validation.Validation{}
	valid.Min(articleID, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	commentService := comment_service.Comment{ArticleID: articleID}
	comments, err := commentService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, comments)
}

func AddComment(c *gin.Context) {
	appG := app.Gin{C: c}

	var postData comment_service.Comment
	err := c.BindJSON(&postData)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	err = postData.Add()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DelComment(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	comment := comment_service.Comment{
		ID: id,
	}
	err := comment.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
