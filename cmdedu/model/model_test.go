package model

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	// GetUser("shineyork").ToString()

	// userDatas = make(map[string]Model, 0)
	// fdata("user", userDataKey, userDatas)
}

func TestGetAdmin(t *testing.T) {
	// GetAdmin("1").ToString()
}
func TestRef(t *testing.T) {
	// ref()
}

func TestFwrite(t *testing.T) {
	user := NewUser()
	user.SetAge(99)
	user.SetSex("男")
	user.SetPassword("123456")
	user.SetUsername("ppp")
	user.Save()
}

func TestFread(t *testing.T) {
	// userDatas = make(map[string]Model, 0)
	// fdata("user", userDataKey, userDatas)
	// for k, v := range userDatas {
	// 	fmt.Println(v)
	// 	fmt.Println(k)
	// 	user, _ := v.(*User)
	// 	fmt.Println(user.GetAge())
	// 	v.ToString()
	// }
}
