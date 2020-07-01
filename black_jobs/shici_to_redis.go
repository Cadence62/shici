package main

import (
	"encoding/json"
	"shici/allmodel"
	"shici/config"
	"shici/utils"
)

func ShiciToRedis() {
	redis_client := utils.Redisclient()

	var shi allmodel.Shi
	allmodel.DB.Last(&shi)
	shi_json, _ := json.Marshal(shi)

	var ci allmodel.Ci
	allmodel.DB.Last(&ci)
	ci_json, _ := json.Marshal(ci)

	var yan allmodel.Yan
	allmodel.DB.Last(&yan)
	yan_json, _ := json.Marshal(yan)

	redis_client.HSet(config.RedisKeyTodayOne, "shi", shi_json)
	redis_client.HSet(config.RedisKeyTodayOne, "ci", ci_json)
	redis_client.HSet(config.RedisKeyTodayOne, "yan", yan_json)
}

func main() {
	allmodel.InitDB()
	ShiciToRedis()
	//ticker:=time.NewTicker(time.Hour*5)
	//go func() {
	//	for _=range ticker.C {
	//		println("test")
	//	}
	//}()
	//
	//select {}
}
