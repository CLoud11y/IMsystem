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
	fmt.Println("PreHandle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
