package app

import (
	"server/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}

func (g *Gin) ResponseList(httpCode, errCode, total int, list interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code":  errCode,
		"msg":   e.GetMsg(errCode),
		"total": total,
		"list":  list,
	})
	return
}
