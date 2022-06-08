package controller

import (
	"github.com/gin-gonic/gin"
	"osier/boot"
)

type Controller struct {

	// 需继承全局控制器
	boot.Ctr
}

// 例子
func (that *Controller) Index(c *gin.Context) {
	that.Suc(c, "你好，Osier框架！", "data")
}
