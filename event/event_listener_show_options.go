package event

import (
	"go-ratel/command"
	"os"
	"strconv"
	"strings"
)

func ListenerShowOptions(ctx *Context, data string) {
	command.PrintNotice("选项: ")
	command.PrintNotice("1. 对战")
	command.PrintNotice("2. 人机")
	command.PrintNotice("3. 设置")
	command.PrintNotice("输入序号进入相应的选项（[EXIT]退出）")

	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("选项")))
	if line == "EXIT" {
		os.Exit(0)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ListenerShowOptionsPVP(ctx, data)
		case 2:
			ListenerShowOptionsPVE(ctx, data)
		case 3:
			ListenerShowOptionsSettings(ctx, data)
		default:
			command.PrintNotice("错误的选项，请重新输入：")
			ListenerShowOptions(ctx, data)
		}
	}
}
