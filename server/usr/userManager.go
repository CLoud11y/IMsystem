package usr

import (
	"fmt"
	"sync"
)

type userManager struct {
	onlineMap map[string]*User
	mapLock   sync.RWMutex
}

var UserManager *userManager

func init() {
	UserManager = &userManager{
		onlineMap: map[string]*User{},
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
