package api_func

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"shici/allmodel"
	"shici/config"
	"shici/utils"
)

func ShiToday(c *gin.Context) {
	redis_client := utils.Redisclient()
	shi_json := redis_client.HGet(config.RedisKeyTodayOne, "shi").Val()

	var shi allmodel.Shi
	json.Unmarshal([]byte(shi_json), &shi)

	result := NewResult()
	result.Code = OK
	result.Data = shi
	c.JSON(200, result)
}

func CiToday(c *gin.Context) {
	redis_client := utils.Redisclient()
	ci_json := redis_client.HGet(config.RedisKeyTodayOne, "ci").Val()

	var ci allmodel.Ci
	json.Unmarshal([]byte(ci_json), &ci)

	result := NewResult()
	result.Code = OK
	result.Data = ci
	c.JSON(200, result)
}

func YanToday(c *gin.Context) {
	redis_client := utils.Redisclient()
	yan_json := redis_client.HGet(config.RedisKeyTodayOne, "yan").Val()

	var yan allmodel.Yan
	json.Unmarshal([]byte(yan_json), &yan)

	result := NewResult()
	result.Code = OK
	result.Data = yan
	c.JSON(200, result)
}
