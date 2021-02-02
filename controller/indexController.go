package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv25/global"
)

type IndexController struct{}
func NewIndexController() IndexController {
	return IndexController{}
}
//得到一件商品的详情
func (g *IndexController) Index(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("success");
	return
}
