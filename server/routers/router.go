package routers

import (
	"net/http"
	_ "server/docs"
	"server/middleware/jwt"
	"server/pkg/qrcode"
	"server/pkg/setting"
	"server/pkg/upload"
	"server/routers/api"
	v1 "server/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	// 获取token
	r.GET("/auth", api.GetAuth)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 添加标签
		apiV1.POST("/tags", v1.AddTag)
		// 修改标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		// 发布文章
		apiV1.POST("/article", v1.AddArticle)
		// 编辑文章
		apiV1.PUT("/article/:id", v1.EditArticle)
		// 删除文章
		apiV1.DELETE("/article/:id", v1.DeleteArticle)

		// 获取某一篇文章的评论列表
		apiV1.POST("/comments", v1.Comments)
		// 发布评论
		apiV1.POST("/comment", v1.AddComment)
		// 删除评论
		apiV1.DELETE("/comment/:id", v1.DelComment)

		apiV1.POST("/articles/poster/generateArticlePoster", v1.GenerateArticlePoster)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	return r
}
