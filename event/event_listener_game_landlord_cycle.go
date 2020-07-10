package event

import "go-ratel/command"

func ListenerGameLandlordCycle(ctx *Context, data string) {
	command.PrintNotice("没有人叫地主，重新发牌。")
}
