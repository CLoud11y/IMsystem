package conf

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConf gorm.Dialector
var GormConf gorm.Option

func init() {
	//配置MySQL连接参数
	username := "root"               //账号
	password := os.Getenv("MYSQLPW") //密码
	host := "8.130.86.24"            //数据库地址，可以是Ip或者域名
	port := 3306                     //数据库端口
	Dbname := "IMsystem"             //数据库名

	if password == "" {
		fmt.Println("the database password is empty!")
	}
	//拼接dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	fmt.Println("DSN:", dsn)

	MysqlConf = mysql.New(mysql.Config{
		DSN: dsn,
	})

	GormConf = &gorm.Config{}
}
