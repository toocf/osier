package router

// 实现路由器的锚点
// 全局中间件实现条件:
// 1.实现 Router 结构体
// 2.方法可导出
// 3.参数是 *gin.Engine
// ::可在同级其他文件中实现
type Router struct{}

/*
// 示例
func (Router) Custom(r *gin.Engine) {
	custom := r.Group("/custom").Use("引入中间件")
	{
		custom.GET("/", func(c *gin.Context) {
			c.String(200, "你好接口")
		})
	}
}
*/
