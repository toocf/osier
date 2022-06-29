package router

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller/admin"
)

func (Router) Admin(r *gin.Engine) {

	// 后台管理端接口
	base := r.Group("/admin")
	{
		// 实例化
		ctr := new(admin.Home)

		// 示例
		base.GET("/", ctr.Index)
	}
}
