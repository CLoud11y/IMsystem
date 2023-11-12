package usr

import (
	"IM_System/conf"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type userManager struct {
	onlineMap map[string]*User
	mapLock   sync.RWMutex

	db *gorm.DB
}

var UserManager *userManager

func init() {
	// 初始化数据库连接
	DB, err := gorm.Open(conf.MysqlConf, conf.GormConf)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{})

	UserManager = &userManager{
		onlineMap: map[string]*User{},
		db:        DB,
	}
}

func (um *userManager) AddUser(user *User) {
	um.mapLock.Lock()
	defer um.mapLock.Unlock()

	_, ok := um.onlineMap[user.Name]
	if ok {
		fmt.Println("addUser fail: userName exist")
		return
	}
	um.onlineMap[user.Name] = user
}

func (um *userManager) RemoveUser(user *User) {
	um.mapLock.Lock()
	defer um.mapLock.Unlock()

	delete(um.onlineMap, user.Name)
}

func (um *userManager) GetUserByName(name string) (*User, bool) {
	um.mapLock.RLock()
	defer um.mapLock.RUnlock()
	user, ok := um.onlineMap[name]
	return user, ok
}

func (um *userManager) GetAllUsers() []*User {
	um.mapLock.RLock()
	defer um.mapLock.RUnlock()
	var users []*User
	for _, each := range um.onlineMap {
		users = append(users, each)
	}
	return users
}

func (um *userManager) RenameUser(user *User, newName string) bool {
	um.mapLock.Lock()
	defer um.mapLock.Unlock()

	_, ok := um.onlineMap[newName]
	if ok {
		// 新用户名已存在 修改失败
		return false
	}

	// 修改用户名
	delete(um.onlineMap, user.Name)
	user.Name = newName
	um.onlineMap[user.Name] = user
	return true
}
