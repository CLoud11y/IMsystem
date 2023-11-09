package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PingRouter MsgIdPing路由
type WhoRouter struct {
	znet.BaseRouter
}

func (r *WhoRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("PreHandle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}

func (r *WhoRouter) Handle(request ziface.IRequest) {
	//查看在线用户
	users := usr.UserManager.GetAllUsers()
	msg := "online users:\n"
	for _, each := range users {
		msg += each.Name + "\n"
	}
	request.GetConnection().SendMsg(common.MsgIdWho, []byte(msg))
}
