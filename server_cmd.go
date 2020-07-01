package main

import (
	"github.com/gin-gonic/gin"
	"shici/allmodel"
	"shici/api_func"
)

func main() {
	//项目启动前初始化
	allmodel.InitDB()

	//web启动
	r := gin.Default()
	// 倒入路由
	api_func.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
