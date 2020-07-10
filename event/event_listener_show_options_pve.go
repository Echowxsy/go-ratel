package event

import (
	"go-ratel/command"
	"strconv"
	"strings"
)

func ListenerShowOptionsPVE(ctx *Context, data string) {
	command.PrintNotice("PVE: ")
	command.PrintNotice("1. Simple Model")
	command.PrintNotice("2. Medium Model")
	command.PrintNotice("3. Difficulty Model")
	command.PrintNotice("输入序号进入相应菜单或输入[BACK|B]返回上级菜单")

	line := command.Write("pve")
	if strings.ToUpper(line) == "BACK"||strings.ToUpper(line) == "B" {
		ListenerShowOptions(ctx, data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		if choose > 0 && choose < 4 {
			ctx.InitLastSellInfo()
			ctx.pushToServer(SERVER_CODE_ROOM_CREATE_PVE, strconv.Itoa(choose))
		} else {
			command.PrintNotice("Invalid option, please choose again：")
			ListenerShowOptionsPVE(ctx, data)
		}
	}
}
