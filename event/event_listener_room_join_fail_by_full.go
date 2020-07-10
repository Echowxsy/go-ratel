package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomJoinFailByFull(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("加入房间失败，房间 " + strconv.Itoa(int(dataMap["roomId"].(float64))) + " 人已经满。")
	ListenerShowOptions(ctx, data)
}
