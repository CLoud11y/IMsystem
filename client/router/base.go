package router

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type MyBaseRouter struct {
	znet.BaseRouter
}

func (r *MyBaseRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("PreHandle: recv from server : msgId=", request.GetMsgID(), ", data:\n", string(request.GetData()))
}
