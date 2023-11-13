package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"

	"github.com/aceld/zinx/ziface"
)

// PingRouter MsgIdPing路由
type PublicRouter struct {
	AuthRouter
}

func (r *PublicRouter) Handle(request ziface.IRequest) {
	//获得要发送的信息
	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	msg := "[public] " + user.Name + ": " + string(request.GetData())
	//将信息发送给全部在线用户
	users := usr.UserManager.GetAllOnlineUsers()
	for _, each := range users {
		each.Conn.SendMsg(common.MsgIdShow, []byte(msg))
	}
}
