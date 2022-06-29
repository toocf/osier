package router

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller/api"
)

func (Router) Api(r *gin.Engine) {

	// 接口路由
	base := r.Group("/api")
	{
		// 实例化
		ctr := new(api.Example)

		base.GET("/", ctr.Index)
	}
}
