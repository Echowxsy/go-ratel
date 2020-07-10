package event

import "go-ratel/command"

func ListenerGamePokerPlayOrderError(ctx *Context, data string) {
	command.PrintNotice("现在不是你的回合，请等待其他玩家操作！")
}
