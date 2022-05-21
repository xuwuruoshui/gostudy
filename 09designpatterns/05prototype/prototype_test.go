package prototype

import (
	"fmt"
	"math/rand"
	"testing"
)

type User struct {
	id int
	name string
}

func (u *User) Clone()Cloneable{
	// 值拷贝一次
	u1 := *u
	return &u1
}

func TestPrototype(t *testing.T){
	
	// 1、创建管理器
	manager := NewPrototypeManager()
	manager.Set("user",&User{
		id: rand.Int(),
		name: "张三",
	})
	u1 := manager.Get("user")
	
	// 2、克隆一次,局部变量相同，但地址不同
	u2 := u1.Clone()
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(&u1,&u2)
	fmt.Println(u1==u2)
}