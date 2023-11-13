package router

import (
	"IM_System/common"
	"IM_System/server/usr"
	"fmt"

	"github.com/aceld/zinx/ziface"
)

// PingRouter MsgIdPing路由
type RenameRouter struct {
	AuthRouter
}

func (r *RenameRouter) Handle(request ziface.IRequest) {
	newName := string(request.GetData())
	u, err := request.GetConnection().GetProperty("user")
	if err != nil {
		fmt.Println("conn.getProperty err: ", err)
	}
	user := u.(*usr.User)
	//查看在线用户
	ok := usr.UserManager.RenameUser(user, newName)
	msg := ""
	if !ok {
		// 已存在该用户名 修改失败
		msg = newName + " already exist"
	} else {
		// 用户名不存在 可以修改
		msg = "your name is changed to: " + newName
	}

	request.GetConnection().SendMsg(common.MsgIdShow, []byte(msg))
}
