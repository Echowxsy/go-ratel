package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerClientNicknameSet(ctx *Context, data string) {
	if data == "" {
		dataMap := make(map[string]interface{})
		if dataMap["invalidLength"] != nil {
			command.PrintNotice("你的昵称长度: " + strconv.Itoa(dataMap["invalidLength"].(int))+" ，太长了。")
		}
	}

	command.PrintNotice("请设置你的昵称（最大长度为：" + strconv.Itoa(NICKNAME_MAX_LENGTH) + " 个字符）")
	nickname := command.DeletePreAndSufSpace(command.Write("nickname"))
	if len(nickname) > NICKNAME_MAX_LENGTH {
		result := make(map[string]interface{})
		result["invalidLength"] = len(nickname)
		resultJson, _ := json.Marshal(&result)
		ListenerClientNicknameSet(ctx, string(resultJson))
	} else {
		ctx.pushToServer(SERVER_CODE_CLIENT_NICKNAME_SET, nickname)
	}
}
