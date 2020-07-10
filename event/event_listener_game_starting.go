package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameStarting(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("游戏开始～")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &pokers)

	command.PrintNotice("")
	command.PrintNotice("\033[32m你的牌是：\033[0m")
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	ListenerGameLandlordElect(ctx, data)
}
