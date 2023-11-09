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
	server.Serve()
}

func DoConnectionBegin(conn ziface.IConnection) {
	user := usr.NewUser(conn.RemoteAddrString(), conn)
	conn.SetProperty("user", user)
	// 先addUser再broadcast 这样用户自己也能收到广播
	usr.UserManager.AddUser(user)
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
	usr.UserManager.RemoveUser(u)
	fmt.Println(u.Name + " is offline")
	go Broadcast(u.Name + " is offline")
}

func Broadcast(msg string) {
	users := usr.UserManager.GetAllUsers()
	for _, each := range users {
		each.Conn.SendMsg(common.MsgIdBroadcast, []byte(msg))
	}
}