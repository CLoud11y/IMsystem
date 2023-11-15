package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConf gorm.Dialector
var GormConf gorm.Option

func init() {
	err := godotenv.Load("conf/db.env")
	if err != nil {
		fmt.Println("load db.env err:", err)
	}

	//配置MySQL连接参数
	username := os.Getenv("MYSQLNAME") //账号
	password := os.Getenv("MYSQLPW")   //密码
	host := os.Getenv("HOST")          //数据库地址，可以是Ip或者域名
	port := os.Getenv("PORT")          //数据库端口
	Dbname := os.Getenv("DBNAME")      //数据库名

	if password == "" {
		fmt.Println("the database password is empty!")
	}
	//拼接dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	fmt.Println("DSN:", dsn)

	MysqlConf = mysql.New(mysql.Config{
		DSN: dsn,
	})

	GormConf = &gorm.Config{}
}
