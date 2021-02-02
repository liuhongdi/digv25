package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv25/global"
	"github.com/liuhongdi/digv25/pkg/sign"
	"regexp"
)
//限流器
func SignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        //检查是否排除当前url
		exclude := regexp.MustCompile("/static/*")
		path := c.Request.RequestURI
		//fmt.Println("path:"+path)
		if exclude != nil && exclude.MatchString(path) {
			c.Next()
			return
		}

		//从header获取参数
		appId := c.Request.Header.Get("appId")
		timestamp := c.Request.Header.Get("timestamp")
		signcli := c.Request.Header.Get("sign")
		nonce := c.Request.Header.Get("nonce")

		//验证签名是否正确
		err := sign.VerifySign(appId,timestamp,signcli,nonce,global.SignAppId,global.SignAppSecret,global.SignApiVersion,global.SignExpire)
        if (err != nil) {
			resultRes := global.NewResult(c)
			resultRes.Error(2004,err.Error())
			return
		} else {
		   //fmt.Println("sign passed")
		   c.Next()
		   return
	    }
	}
}


