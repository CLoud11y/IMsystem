package router

import (
	"IM_System/common"

	"github.com/aceld/zinx/ziface"
)

// PingRouter MsgIdPing路由
type PingRouter struct {
	MyBaseRouter
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	request.GetConnection().SendMsg(common.MsgIdShow, []byte("pong...pong...pong...[FromServer]"))
}
