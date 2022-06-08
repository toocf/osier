package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 实现全局中间件的锚点
// 全局中间件实现条件:
// 1.实现 Middle 结构体
// 2.方法可导出
// 3.参数是 *gin.Context
// ::可在同级其他文件中实现
type Middle struct{}

// 示例
func (Middle) Custom(c *gin.Context) {
	fmt.Println("自定义全局中间件Custom执行！")
}
