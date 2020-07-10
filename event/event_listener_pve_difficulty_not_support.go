package event

import "go-ratel/command"

func ListenerPVEDifficultyNotSupport(ctx *Context, data string) {
	command.PrintNotice("现在不支持你选择的难度。\n")
	ListenerShowOptionsPVE(ctx, data)
}
