package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.listen() err:", err)
		return
	}
	// close listen socket
	defer listener.Close()
	// 开启监听message
	go server.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// do handler
		go server.Handler(conn)
	}
}

func (server *Server) Handler(conn net.Conn) {
	user := NewUser(conn)
	//用户上线 将用户加入OnlineMap中
	user.Online(server)
	//监听用户是否活跃的channel 不活跃被强踢
	isActive := make(chan bool)

	//接收客户端消息并处理
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				isActive <- false
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			//
			msg := string(buf[:n])
			user.DoMessage(msg, server)
			//用户发送消息 视为活跃状态
			isActive <- true
		}
	}()

	//超时强踢
	timeout := 100
	for {
		select {
		case flag := <-isActive:
			if flag {
				// do nothing 相当于重置计时器
			} else {
				user.Offline(server)
				return
			}
		case <-time.After(time.Second * time.Duration(timeout)):
			user.C <- "you are offline for inactivity"
			user.Offline(server)
			return
		}
	}
}

func (server *Server) Brodcast(user *User, msg string) {
	sendMsg := "[" + user.Name + "]" + " to all: " + msg
	server.Message <- sendMsg
}

func (server *Server) ListenMessage() {
	for {
		msg := <-server.Message
		server.mapLock.Lock()
		for _, client := range server.OnlineMap {
			client.C <- msg
		}
		server.mapLock.Unlock()
	}
}
