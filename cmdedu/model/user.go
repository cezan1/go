package model

import (
	"strconv"
)

type User struct {
	username string
	password string
	age      int
	sex      string
}

// 根据用户端 key 为 username ，value 为 User模型
var (
	userDatas   map[string]Model // 用于存放之后user查询的信息
	userDataKey string           = "username"
)

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}
func (u *User) SetAge(age int) {
	u.age = age
}
func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetAge() int {
	return u.age
}
func (u *User) GetSex() string {
	return u.sex
}
func (u *User) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

// 用于保存模型信息
func (u *User) Save() bool {
	// 先在内存里面保存一份
	userDatas[u.username] = u
	return fwrite("user", userDatas)
}

func (u *User) All() []*User {
	var users []*User = make([]*User, 0)
	// fmt.Println("容量", len(userDatas))
	// return userDatas
	for _, user := range userDatas {
		users = append(users, user.(*User))
		// fmt.Println("每次的变化", len(users))
	}
	// fmt.Println("整顿后端user", len(users))
	return users
}

func GetUser(username string) *User {
	return userDatas[username].(*User)
}
