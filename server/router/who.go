package router

import (
	"IM_System/common"
	"IM_System/server/usr"

	"github.com/aceld/zinx/ziface"
)

// PingRouter MsgIdPing路由
type WhoRouter struct {
	MyBaseRouter
}

func (r *WhoRouter) Handle(request ziface.IRequest) {
	//查看在线用户
	users := usr.UserManager.GetAllUsers()
	msg := "online users:\n"
	for _, each := range users {
		msg += each.Name + "\n"
	}
	request.GetConnection().SendMsg(common.MsgIdShow, []byte(msg))
}
