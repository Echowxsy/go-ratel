package event

import (
	"encoding/json"
	"fmt"
	"go-ratel/command"
	"strconv"
	"strings"
)

var choose = [...]string{"UP", "DOWN"}
var positionString = [...]string{"上家", "下家"}
var format = "\n[%-2s] %-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s [%-2s] 剩余牌数： %-2s"

func ListenerGamePokerPlayRedirect(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	sellClientId := int(dataMap["sellClientId"].(float64))

	clientInfos := make([]map[string]interface{}, 0)
	clientInfoBytes, _ := json.Marshal(dataMap["clientInfos"])
	_ = json.Unmarshal(clientInfoBytes, &clientInfos)

	for index := 0; index < 2; index++ {
		for _, clientInfo := range clientInfos {
			position := clientInfo["position"].(string)
			if strings.ToUpper(position) == strings.ToUpper(choose[index]) {
				var roleString string
				if ( clientInfo["type"].(string)=="PEASANT") {
					roleString = "农民"
				} else {
					roleString = "地主"
				}
				command.PrintNotice(fmt.Sprintf(format, positionString[index], clientInfo["clientNickname"].(string),roleString, strconv.Itoa(int(clientInfo["surplus"].(float64)))))
			}
		}
	}
	command.PrintNotice("")
	if sellClientId == ctx.UserId {
		ListenerGamePokerPlay(ctx, data)
	} else {
		command.PrintNotice("现在出牌等人是玩家是 " + dataMap["sellClinetNickname"].(string) + " ，请等 TA 出牌。")
	}
}
