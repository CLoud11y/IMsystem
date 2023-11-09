package usr

import "github.com/aceld/zinx/ziface"

type User struct {
	Name string
	Conn ziface.IConnection
}

func NewUser(name string, conn ziface.IConnection) *User {
	return &User{
		Name: name,
		Conn: conn,
	}
}
