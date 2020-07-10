package event

import (
	"go-ratel/command"
	"strconv"
	"strings"
)

func ListenerShowOptionsPVP(ctx *Context, data string) {
	command.PrintNotice("对战: ")
	command.PrintNotice("1. 创建房间")
	command.PrintNotice("2. 查看房间")
	command.PrintNotice("3. 加入房间")
	command.PrintNotice("输入序号进入相应菜单或输入[BACK|B]返回上级菜单")
	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("对战")))
	if line == "BACK"||line == "B" {
		ListenerShowOptions(ctx, data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ctx.pushToServer(SERVER_CODE_ROOM_CREATE, "")
		case 2:
			ctx.pushToServer(SERVER_CODE_GET_ROOMS, "")
		case 3:
			command.PrintNotice("输入房间ID加入房间或输入[BACK|B]返回上级菜单：")
			line := command.DeletePreAndSufSpace(command.Write("房间ID"))
			if strings.ToUpper(line) == "BACK"||strings.ToUpper(line) == "B" {
				ListenerShowOptionsPVP(ctx, data)
			} else {
				roomid, e := strconv.Atoi(line)
				if e != nil {
					roomid = -1
				}
				if roomid < 1 {
					command.PrintNotice("\033[31m错误的输入，请重新输入：\033[0m")
					ListenerShowOptionsPVP(ctx, data)
				} else {
					ctx.pushToServer(SERVER_CODE_ROOM_JOIN, strconv.Itoa(roomid))
				}
			}
		default:
			command.PrintNotice("\033[31m错误的输入，请重新输入：\033[0m")
			ListenerShowOptionsPVP(ctx, data)
		}
	}
}
