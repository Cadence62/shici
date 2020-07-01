package api_func

import (
	"github.com/gin-gonic/gin"
)


func PingExample(c *gin.Context){
	result := NewResult()
	result.Code =OK
	result.Data = "Ping Example"
	c.JSON(200,result)
}