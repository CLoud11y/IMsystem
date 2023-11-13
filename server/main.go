package main

import (
	"IM_System/common"
	"IM_System/server/router"
	"IM_System/server/usr"
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

func main() {
	server := znet.NewServer()

	server.SetOnConnStart(DoConnectionBegin)
	server.SetOnConnStop(DoConnectionLost)

	server.AddRouter(common.MsgIdPing, &router.PingRouter{})
	server.AddRouter(common.MsgIdWho, &router.WhoRouter{})
	server.AddRouter(common.MsgIdRename, &router.RenameRouter{})
	server.AddRouter(common.MsgIdPublic, &router.PublicRouter{})
	server.AddRouter(common.MsgIdPrivate, &router.PrivateRouter{})
	server.AddRouter(common.MsgIdLogin, &router.LoginRouter{})

	server.Serve()
}

func DoConnectionBegin(conn ziface.IConnection) {
	user := usr.NewUser(conn.RemoteAddrString(), conn, "", false)
	conn.SetProperty("user", user)
	// 先addUser再broadcast 这样用户自己也能收到广播
	usr.UserManager.AddOnlineUser(user)
	fmt.Println(user.Name + " is online")
	go Broadcast(user.Name + " is online")
}

func DoConnectionLost(conn ziface.IConnection) {
	user, err := conn.GetProperty("user")
	if err != nil {
		fmt.Println("conn.GetProperty err: ", err)
	}
	//TODO
	u := user.(*usr.User)
	usr.UserManager.RemoveOnlineUser(u)
	fmt.Println(u.Name + " is offline")
	go Broadcast(u.Name + " is offline")
}

func Broadcast(msg string) {
	users := usr.UserManager.GetAllOnlineUsers()
	for _, each := range users {
		each.Conn.SendMsg(common.MsgIdShow, []byte(msg))
	}
}
