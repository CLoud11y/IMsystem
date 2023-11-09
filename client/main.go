package main

import (
	"IM_System/client/router"
	"IM_System/common"
	"fmt"
	"strings"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// 客户端自定义业务
func waitingInput(conn ziface.IConnection) {
	for {
		var input string
		n, _ := fmt.Scanln(&input)
		if n == 0 {
			continue
		}
		splited := strings.Split(input, "|")
		instruction := splited[0]
		msgId, ok := common.InstructionMap[instruction]
		if !ok {
			continue
		}
		err := conn.SendMsg(msgId, []byte(input))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

// 创建连接的时候执行
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go waitingInput(conn)
}

func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 8888)

	//设置链接建立成功后的钩子函数
	client.SetOnConnStart(onClientStart)

	client.AddRouter(common.MsgIdPong, &router.PongRouter{})
	client.AddRouter(common.MsgIdWho, &router.WhoRouter{})
	client.AddRouter(common.MsgIdBroadcast, &router.MyBaseRouter{})

	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}