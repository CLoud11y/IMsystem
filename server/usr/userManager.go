package usr

import (
	"IM_System/conf"
	"IM_System/utils"
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

func (um *userManager) Login(user *User, userName string, password string) string {
	//检查在线用户中是否已经存在
	if _, ok := um.GetOnlineUserByName(userName); ok {
		return "login fail:" + userName + " is already online"
	}
	tempUser := &User{}
	err := um.db.Where("name = ?", userName).First(tempUser).Error
	if err != nil {
		msg := "login fail:"
		if err == gorm.ErrRecordNotFound {
			msg += " userName: " + userName + " not found"
		}
		return msg
	}
	// 检查密码
	md5_password := utils.Encrypt(password)
	if md5_password != tempUser.Password {
		return "login fail: password err"
	}

	// 用户名密码正确 成功登录
	um.mapLock.Lock()

	delete(um.onlineMap, user.Name)
	user.Id = tempUser.Id
	user.Name = tempUser.Name
	user.Password = tempUser.Password
	user.Islogin = true
	um.onlineMap[user.Name] = user

	um.mapLock.Unlock()
	return "login success"
}

func (um *userManager) Signup(user *User, userName string, password string) string {
	// 检查用户名是否已存在
	tempUser := &User{}
	err := um.db.Where("name = ?", userName).First(tempUser).Error
	// 用户已存在
	if err == nil {
		return "login fail: userName: " + userName + " exists"
	}
	// 发生未知错误
	if err != gorm.ErrRecordNotFound {
		return err.Error()
	}
	// 正常注册流程
	tempUser.Name = userName
	tempUser.Password = utils.Encrypt(password)
	um.db.Save(&tempUser)

	return "signup success, please login"
}

func (um *userManager) AddOnlineUser(user *User) {
	um.mapLock.Lock()
	defer um.mapLock.Unlock()

	_, ok := um.onlineMap[user.Name]
	if ok {
		fmt.Println("addUser fail: userName exist")
		return
	}
	um.onlineMap[user.Name] = user
}

func (um *userManager) RemoveOnlineUser(user *User) {
	um.mapLock.Lock()
	defer um.mapLock.Unlock()

	delete(um.onlineMap, user.Name)
}

func (um *userManager) GetOnlineUserByName(name string) (*User, bool) {
	um.mapLock.RLock()
	defer um.mapLock.RUnlock()
	user, ok := um.onlineMap[name]
	return user, ok
}

func (um *userManager) GetAllOnlineUsers() []*User {
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

	// 查询数据库中是否存在newname
	if err := um.db.Where("name = ?", newName).First(&User{}).Error; err != gorm.ErrRecordNotFound {
		return false
	}

	// 修改用户名
	delete(um.onlineMap, user.Name)
	user.Name = newName
	um.onlineMap[user.Name] = user
	// 更新数据库
	um.db.Save(&user)

	return true
}
