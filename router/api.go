package router

import (
	"github.com/gin-gonic/gin"
)

func (Router) Api(r *gin.Engine) {

	// 接口路由
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "Gin原生消息",
				"data": "数据",
			})
		})
	}
}
