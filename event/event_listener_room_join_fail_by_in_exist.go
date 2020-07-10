package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerRoomJoinFailByInExist(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("加入房间失败，房间 " + dataMap["roomId"].(string) + " 不存在！")
	ListenerShowOptions(ctx, data)
}
