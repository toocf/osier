package admin

import (
	"github.com/gin-gonic/gin"
	"osier/app/controller"
)

type Index struct {
	controller.Controller
}

// 重写继承的Index
func (that *Index) Index(c *gin.Context) {

	that.Suc(c, "Osier Admin Index!!!", "[2226788556]")
}
