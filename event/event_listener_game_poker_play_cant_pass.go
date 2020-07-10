package event

import "go-ratel/command"

func ListenerGamePokerPlayCantPass(ctx *Context, data string) {
	command.PrintNotice("现在是你出牌的时候，你不可以跳过。")
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
