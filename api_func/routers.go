package api_func

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	urls_api_v1 := route.Group("/api")
	{
		//测试Url
		urls_api_v1.GET("/ping", PingExample)
		//获取进入诗词推送
		urls_api_v1.GET("/shitoday", ShiToday)
		urls_api_v1.GET("/citoday", CiToday)
		urls_api_v1.GET("/yantoday", YanToday)
	}
}
