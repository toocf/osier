package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osier/app/controller"
)

type Index struct {
	controller.Controller
}

// 接口首页示例
// @Summary 这是index接口
// @Schemes
// @Description 这是index接口的简介
// @Tags 接口示例
// @Accept json
// @Produce json
// @Success 200 {string} index
// @Router /api/index [get]
func (that *Index) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
