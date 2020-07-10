package event

import (
	"encoding/json"
	"go-ratel/command"
	"strings"
)

func ListenerGameLandlordElect(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)
	turnClientId := int(dataMap["nextClientId"].(float64))

	if dataMap["preClientNickname"] != nil {
		command.PrintNotice(dataMap["preClientNickname"].(string) + " 不叫地主！")
	}
	if turnClientId == ctx.UserId {
		command.PrintNotice("\033[32m现在的你的回合，你要叫地主吗？[Y/N] 叫/不叫（[EXIT]退出房间)\033[0m")
		line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("Y/N")))
		switch line {
		case "EXIT":
			ctx.pushToServer(SERVER_CODE_CLIENT_EXIT, "")
		case "Y":
			ctx.pushToServer(CODE_GAME_LANDLORD_ELECT, "TRUE")
		case "N":
			ctx.pushToServer(CODE_GAME_LANDLORD_ELECT, "FALSE")
		default:
			command.PrintNotice("\033[31m错误的输入\033[0m")
			ListenerGameLandlordElect(ctx, data)
		}
	} else {
		command.PrintNotice("现在是 " + dataMap["nextClientNickname"].(string) + " 的回合，请等待。")
	}
}
