package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	go user.ListenMessage()
	return user
}

func (user *User) ListenMessage() {
	for msg := range user.C {
		_, err := user.conn.Write([]byte(msg + "\r\n"))
		if err != nil {
			panic(err)
		}
	}
	//不监听后关闭conn，conn在这里关闭最合适
	err := user.conn.Close()
	if err != nil {
		panic(err)
	}
}

func (user *User) Online(server *Server) {
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	server.Brodcast(user, " is online")
}

func (user *User) Offline(server *Server) {
	server.mapLock.Lock()
	delete(server.OnlineMap, user.Name)
	server.mapLock.Unlock()
	server.Brodcast(user, " is offline")
	// 关闭channel，进而结束监听channel的goroutine
	close(user.C)
}

func (user *User) DoMessage(msg string, server *Server) {
	if msg == "who" {
		// 查询当前在线用户
		server.mapLock.Lock()
		onlineMsg := "online users:\r\n"
		for _, each := range server.OnlineMap {
			onlineMsg += "[" + each.Addr + "]" + each.Name + "\r\n"
		}
		server.mapLock.Unlock()
		user.C <- onlineMsg

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 改名 消息格式：rename|<newname>
		newname := msg[7:]
		// 查看newname是否被占用
		_, ok := server.OnlineMap[newname]
		if ok {
			failMsg := "rename fail:" + newname + "is used\r\n"
			user.C <- failMsg
		} else {
			server.mapLock.Lock()
			delete(server.OnlineMap, user.Name)
			server.OnlineMap[newname] = user
			server.mapLock.Unlock()

			user.Name = newname
			succMsg := "your name is changed to " + newname
			user.C <- succMsg
		}

	} else {
		server.Brodcast(user, msg)
	}
}
