package event

import "go-ratel/command"

func ListenerGamePokerPlayInvalid(ctx *Context, data string) {
	command.PrintNotice("\033[31m牌型错误！\033[0m")

	var roleString string
	if (ctx.LastSellClientType=="PEASANT") {
		roleString = "农民"
	} else {
		roleString = "地主"
	}

	if ctx.LastPokers != nil {
		command.PrintNotice(ctx.LastSellClientNickname + "[" + roleString + "] 出了：")
		command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
	}
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
