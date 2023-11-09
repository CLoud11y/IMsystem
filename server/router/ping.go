package router

import (
	"IM_System/common"
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PingRouter MsgIdPing路由
type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("PreHandle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	request.GetConnection().SendMsg(common.MsgIdPong, []byte("pong...pong...pong...[FromServer]"))
}
