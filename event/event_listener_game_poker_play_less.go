package event

import "go-ratel/command"

func ListenerGamePokerPlayLess(ctx *Context, data string) {
	command.PrintNotice("你的牌比上家小，请选择新的牌或者输入[PASS|P]跳过")
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
