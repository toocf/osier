package boot

import (
	"github.com/gin-gonic/gin"
)

// 总控制器
type Ctr struct{}

// 成功返回 - 没有消息就是最好的消息
func (this Ctr) Suc(c *gin.Context, msg string, data any) {
	this.Res(c, msg, data, 0)
}

// 失败返回 - 成功只有一种，而失败却有无数种可能
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
