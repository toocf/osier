package api

import (
	"github.com/gin-gonic/gin"
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

	that.Suc(c, "Osier API Index!!!", "[2226788556]")
}
