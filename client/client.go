package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	mode       int
}

func (client *Client) ShowMenu() bool {
	var mode int
	fmt.Println("1. public chat")
	fmt.Println("2. private chat")
	fmt.Println("3. change name")
	fmt.Println("0. quit")

	fmt.Scanln(&mode)
	if mode >= 0 && mode <= 3 {
		client.mode = mode
		return true
	} else {
		fmt.Println("input value illegal")
		return false
	}
}

func (client *Client) Run() {
	for client.mode != 0 {
		for !client.ShowMenu() {
		}

		switch client.mode {
		case 0:
			return
		case 1:
			client.PublicChat()
		case 2:
			client.PrivateChat()
		case 3:
			client.UpdateName()
		}
	}

}

func (client *Client) ShowUsers() {
	msg := "who"
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	fmt.Println("mode: private(enter 'exit' to exit)")
	client.ShowUsers()

	fmt.Println("enter the user's name:")
	var targetName string
	fmt.Scanln(&targetName)
	if targetName == "exit" {
		return
	}
	var msg string
	for {
		fmt.Printf("to [%s]: ", targetName)
		fmt.Scanln(&msg)
		if msg == "exit" {
			return
		}
		if msg == "" {
			continue
		}
		Sendmsg := "to|" + targetName + "|" + msg
		_, err := client.conn.Write([]byte(Sendmsg))
		if err != nil {
			fmt.Println("conn Write err:", err)
			return
		}
		msg = ""
	}
}

func (client *Client) PublicChat() {
	fmt.Println("mode: public(enter 'exit' to exit)")
	var msg string
	for {
		fmt.Scanln(&msg)
		if msg == "exit" {
			return
		}
		if msg == "" {
			continue
		}
		_, err := client.conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("conn Write err:", err)
			return
		}
		msg = ""
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("Please enter new name:")
	fmt.Scanln(&client.Name)
	msg := "rename|" + client.Name
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 处理服务器发来的消息
func (client *Client) DealResponse() {
	// 一直阻塞监听client.conn 并重定向到stdout
	io.Copy(os.Stdout, client.conn)
}

func NewClient(serverIp string, serverPort int) (*Client, error) {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		mode:       -1,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
	}
	client.conn = conn
	return client, err
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set serverIp")
	flag.IntVar(&serverPort, "port", 8888, "set serverPort")
}

func main() {
	flag.Parse()

	client, err := NewClient(serverIp, serverPort)
	if err != nil {
		fmt.Println("----sever connect fail----")
		return
	}
	fmt.Println("----sever connect succuss----", client)
	go client.DealResponse()
	client.Run()
}
