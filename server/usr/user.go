package usr

import (
	"github.com/aceld/zinx/ziface"
)

type User struct {
	Name     string             `gorm:"primaryKey; not null; type: varchar(32);"`
	Password string             `gorm:"not null; type: char(32);"`
	Conn     ziface.IConnection `gorm:"-"`
}

func NewUser(name string, conn ziface.IConnection, pswd string) *User {
	return &User{
		Name:     name,
		Conn:     conn,
		Password: pswd,
	}
}
