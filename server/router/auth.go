package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// 权限校验 在prehandle中检查用户是否登录 需要权限校验的其他路由继承此类
type AuthRouter struct {
	znet.BaseRouter
}

func (r *AuthRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Printf("AUTH PreHandle: recv from client[%s]: msgId=%d, data=%s\n",
		request.GetConnection().RemoteAddrString(), request.GetMsgID(), string(request.GetData()))

	// 检查用户是否已登录
	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	if !user.Islogin {
		// 未登录 提示用户登录 中断后续handler
		user.Conn.SendMsg(common.MsgIdShow, []byte("please login..."))
		fmt.Println("AUTH fail...")
		request.Abort()
		return
	}
	fmt.Println("AUTH success...")
}
