package routes

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller/admin"
)

func (Route) Admin(r *gin.Engine) {

	// 后台管理端接口
	base := r.Group("/admin")
	{
		// 实例化
		ctr := new(admin.Index)

		// 示例
		base.GET("/", ctr.Index)
	}
}
