package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerShowPokers(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	ctx.LastSellClientNickname = dataMap["clientNickname"].(string)
	ctx.LastSellClientType = dataMap["clientType"].(string)

	var roleString string
	if (ctx.LastSellClientType=="PEASANT") {
		roleString = "农民"
	} else {
		roleString = "地主"
	}
	command.PrintNotice(ctx.LastSellClientNickname + "[" + roleString + "] 出了：")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal(pokersBytes, &pokers)
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	if dataMap["sellClinetNickname"] != nil {
		command.PrintNotice("下一个出牌的人是 " + dataMap["sellClinetNickname"].(string) + "，请等 TA 出牌。")
	}
}
