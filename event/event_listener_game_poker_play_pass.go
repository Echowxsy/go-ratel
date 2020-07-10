package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGamePokerPlayPass(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice(dataMap["clientNickname"].(string) + " 过/要不起，现在是 " + dataMap["nextClientNickname"].(string) + " 的回合。")

	turnClientId := int(dataMap["nextClientId"].(float64))
	if ctx.UserId == turnClientId {
		ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
	}
}
