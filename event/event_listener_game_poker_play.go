package event

import (
	"encoding/json"
	"go-ratel/command"
	"strings"
)

func ListenerGamePokerPlay(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("\033[32m现在是你出牌的时候，你的牌是：\033[0m")


	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &pokers)
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	command.PrintNotice("\033[32m输入你要出的牌([EXIT]退出房间｜[PASS|P]不出牌｜[0]牌10｜[S]小王｜[X]大王)\033[0m")
	line := command.DeletePreAndSufSpace(command.Write("card"))
	if line == "" {
		command.PrintNotice("\033[31m错误的输入\033[0m")
		ListenerGamePokerPlay(ctx, data)
	} else {
		if strings.ToUpper(line) == "PASS" || strings.ToUpper(line) == "P"{
			ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_PASS, "")
		} else if strings.ToUpper(line) == "EXIT" {
			ctx.pushToServer(SERVER_CODE_CLIENT_EXIT, "")
		} else {
			strs := strings.Split(line, " ")
			options := make([]string, 0)
			access := true
			for i := 0; i < len(strs); i++ {
				str := strs[i]
				for _, v := range []byte(str) {
					if string(v) == " " || string(v) == "\t" {
					} else {
						if !pokerLevelAliasContainer(v) {
							access = false
							break
						} else {
							options = append(options, string(v))
						}
					}
				}
			}
			if access {
				bytes, _ := json.Marshal(&options)
				ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY, string(bytes))
			} else {
				command.PrintNotice("\033[31m错误的输入\033[0m")
				var roleString string
				if (ctx.LastSellClientType=="PEASANT") {
					roleString = "农民"
				} else {
					roleString = "地主"
				}
				if ctx.LastPokers != nil {
					command.PrintNotice(ctx.LastSellClientNickname + "[" + roleString + "] playd:")
					command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
				}
				ListenerGamePokerPlay(ctx, data)
			}
		}
	}
}

func pokerLevelAliasContainer(b byte) bool {
	pokerAlias := []string{"3", "4", "5", "6", "7", "8", "9", "T", "t", "0", "J", "j", "Q", "q", "K", "k", "A", "a", "1", "2", "S", "s", "X", "x"}
	for _, v := range pokerAlias {
		if v == string(b) {
			return true
		}
	}
	return false
}
