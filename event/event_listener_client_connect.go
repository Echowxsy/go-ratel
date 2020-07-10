package event

import (
	"go-ratel/command"
	"strconv"
)

func ListenerClientConnect(ctx *Context, data string) {
	command.PrintNotice("链接服务器成功，欢迎进入Ratel！")
	ctx.UserId, _ = strconv.Atoi(data)
}
