package routes

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller/home"
)

func (Route) Home(r *gin.Engine) {

	// 加载模板
	r.LoadHTMLGlob("resources/views/home/*")
	// 接口路由
	base := r.Group("/")
	{
		// 实例化
		ctr := new(home.Index)

		base.GET("/", ctr.Index)
	}
}
