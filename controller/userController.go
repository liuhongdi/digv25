package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv25/global"
)

type UserController struct{}
func NewUserController() UserController {
	return UserController{}
}
//得到一件商品的详情
func (g *UserController) Login(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("success");
	return
}