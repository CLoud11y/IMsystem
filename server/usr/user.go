package usr

import (
	"github.com/aceld/zinx/ziface"
)

type User struct {
	Id       int                `gorm:"primaryKey; not null;"`
	Name     string             `gorm:"unique; not null; type: varchar(32);"`
	Password string             `gorm:"not null; type: varchar(32);"`
	Conn     ziface.IConnection `gorm:"-"`
	Islogin  bool               `gorm:"-"`
}

func NewUser(name string, conn ziface.IConnection, pswd string, isLogin bool) *User {
	return &User{
		Name:     name,
		Conn:     conn,
		Password: pswd,
		Islogin:  isLogin,
	}
}
