package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv25/controller"
	"github.com/liuhongdi/digv25/global"
	"github.com/liuhongdi/digv25/middleware"
	"log"
	"net/http"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(middleware.SignMiddleware())
	router.Use(Recover)
	//static
	router.StaticFS("/static", http.Dir("/data/liuhongdi/digv25/static"))
	// 路径映射:index
	indexc:=controller.NewIndexController()
	router.GET("/index/index", indexc.Index);
	// 路径映射:user
	userc:=controller.NewUserController()
	router.POST("/user/login", userc.Login);
	return router
}

func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404,"资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}