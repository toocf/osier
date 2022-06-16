package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 授权登录
func AuthLogin(c *gin.Context) {

	fmt.Println("授权登录中间件，需结合用户表制定功能！")
}
