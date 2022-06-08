package router

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller"
)

func (Router) Admin(r *gin.Engine) {

	// 后台管理端接口
	admin := r.Group("/admin")
	{
		// 实例化
		clr := new(controller.Controller)

		// 示例
		admin.GET("/", clr.Index)
	}
}
