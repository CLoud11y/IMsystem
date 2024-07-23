package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"
	"strings"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type SignupRouter struct {
	znet.BaseRouter
}

func (r *SignupRouter) PreHandle(request ziface.IRequest) {
	fmt.Printf("Signup PreHandle: recv from client[%s]: msgId=%d, data=%s\n",
		request.GetConnection().RemoteAddrString(), request.GetMsgID(), string(request.GetData()))

	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	if user.Islogin {
		// 该用户已登录 中断后续handle
		user.Conn.SendMsg(common.MsgIdShow, []byte("you already logged in, userName: "+user.Name))
		request.Abort()
	}
}

// 重写登录路由的handle 调用userManager的注册接口
func (r *SignupRouter) Handle(request ziface.IRequest) {
	//获得注册信息 格式为userName|password
	data := string(request.GetData())
	temp := strings.Split(data, "|")
	userName := temp[0]
	password := ""
	if len(temp) > 1 {
		password = temp[1]
	}
	//获取当前用户
	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	msg := usr.UserManager.Signup(user, userName, password)
	user.Conn.SendMsg(common.MsgIdShow, []byte(msg))
}
