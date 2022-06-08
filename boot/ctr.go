package boot

import (
	"github.com/gin-gonic/gin"
)

// 总控制器
type Ctr struct{}

// 成功返回
func (this Ctr) Suc(c *gin.Context, msg string, data any) {
	this.Res(c, msg, data, 0)
}

// 失败返回
func (this Ctr) Err(c *gin.Context, msg string, data any) {
	this.Res(c, msg, data, 1)
}

// 普通返回
func (Ctr) Res(c *gin.Context, msg string, data any, code int) {
	Res(c, msg, data, code)
}

// 通用返回
func Res(c *gin.Context, msg string, data any, code int) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  Lang(msg),
	})
}
