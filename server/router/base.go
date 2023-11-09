package router

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PingRouter MsgIdPing路由
type MyBaseRouter struct {
	znet.BaseRouter
}

func (r *MyBaseRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Printf("PreHandle: recv from client[%s]: msgId=%d, data=%s\n", request.GetConnection().RemoteAddrString(), request.GetMsgID(), string(request.GetData()))
}
