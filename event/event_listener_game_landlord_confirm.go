package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameLandlordConfirm(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	landlordNickname := dataMap["landlordNickname"].(string)
	command.PrintNotice(landlordNickname + " 叫地主，获得额外的三张牌：")

	// 序列化
	additionalPokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["additionalPokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &additionalPokers)

	command.PrintPokers(additionalPokers, ctx.PokerPrinterType)
	ctx.pushToServer(CODE_GAME_POKER_PLAY_REDIRECT, "")
}
