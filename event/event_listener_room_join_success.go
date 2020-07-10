package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomJoinSuccess(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	ctx.InitLastSellInfo()
	joinClientId := int(dataMap["clientId"].(float64))

	if ctx.UserId == joinClientId {
		command.PrintNotice("你加入了房间：" + strconv.Itoa(int(dataMap["roomId"].(float64))) + "。现在有 " + strconv.Itoa(int(dataMap["roomClientCount"].(float64))) + " 个玩家。")
		command.PrintNotice("请等待玩家加入，当有三个玩家的时候游戏会自动开始。")
	} else {
		command.PrintNotice(dataMap["clientNickname"].(string) + " 加入了房间，现在有 " + strconv.Itoa(int(dataMap["roomClientCount"].(float64)))+ " 个玩家。")
	}
}
