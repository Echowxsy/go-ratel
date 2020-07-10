package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomCreateSuccess(ctx *Context, data string) {
	room := Room{}
	_ = json.Unmarshal([]byte(data), &room)

	ctx.InitLastSellInfo()

	command.PrintNotice("你创建的房间ID为： " + strconv.Itoa(room.Id))
	command.PrintNotice("请等待其他玩家加入！")
}
