package event

import "go-ratel/command"

func ListenerClientKick(ctx *Context, data string) {
	command.PrintNotice("\033[31m长时间未操作，踢出房间。\033[0m\n")
	ListenerShowOptions(ctx, data)
}
