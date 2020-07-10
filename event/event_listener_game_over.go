package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameOver(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	var roleString string
	if (dataMap["winnerType"].(string)=="PEASANT") {
		roleString = "农民"
	} else {
		roleString = "地主"
	}
	
	command.PrintNotice("\n玩家 " + dataMap["winnerNickname"].(string) + "[" + roleString + "]" + " 赢了")
	command.PrintNotice("游戏结束！\n")
}
