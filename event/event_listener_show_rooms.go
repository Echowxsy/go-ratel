package event

import (
	"encoding/json"
	"fmt"
	"go-ratel/command"
	"strconv"
	"strings"
)

func ListenerShowRooms(ctx *Context, data string) {
	roomList := make([]map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(data), &roomList)

	if len(roomList) > 0 {
		format := "#\t%s\t|\t%-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s\t|\t%-6s\t|\t%-6s\t#"
		command.PrintNotice(fmt.Sprintf(format, "ID", "OWNER", "COUNT", "TYPE"))
		for _, room := range roomList {
			command.PrintNotice(fmt.Sprintf(format, strconv.Itoa(int(room["roomId"].(float64))), room["roomOwner"].(string), strconv.Itoa(int(room["roomClientCount"].(float64))), room["roomType"].(string)))
		}
		command.PrintNotice("")
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
	} else {
		command.PrintNotice("没有可以加入的房间，请创建一个房间！")
		ListenerShowOptionsPVP(ctx, data)
	}
}
