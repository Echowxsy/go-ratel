package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerGamePokerPlayMismatch(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("你的牌型是 " + dataMap["playType"].(string) + " (" + strconv.Itoa(int(dataMap["playCount"].(float64))) +
		") 但是上家的牌型是 " + dataMap["preType"].(string) + " (" + strconv.Itoa(int(dataMap["preCount"].(float64))) + ")，牌型不匹配！")

	var roleString string
	if (ctx.LastSellClientType=="PEASANT") {
		roleString = "农民"
	} else {
		roleString = "地主"
	}
	if ctx.LastPokers != nil {
		command.PrintNotice(ctx.LastSellClientNickname + "[" + roleString + "] 出了：")
		command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
	}
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
