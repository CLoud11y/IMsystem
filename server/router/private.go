package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"
	"strings"

	"github.com/aceld/zinx/ziface"
)

// PingRouter MsgIdPing路由
type PrivateRouter struct {
	MyBaseRouter
}

func (r *PrivateRouter) Handle(request ziface.IRequest) {
	//获得要发送的信息 私聊信息格式为targetName|msg
	data := string(request.GetData())
	temp := strings.Split(data, "|")
	targetName := temp[0]
	content := ""
	if len(temp) > 1 {
		content = temp[1]
	}
	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	msg := "[private] " + user.Name + ": " + content
	//将信息发送给目标用户
	targetUser, ok := usr.UserManager.GetUserByName(targetName)
	if ok {
		targetUser.Conn.SendMsg(common.MsgIdShow, []byte(msg))
	} else {
		user.Conn.SendMsg(common.MsgIdShow, []byte(targetName+" is not online"))
	}
}
