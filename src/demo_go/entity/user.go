package entity

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func (user *User) C(name string, age int, sex string) User {
	return User{name: name, age: age, sex: sex}
}

func (user *User) Say() {
	fmt.Println("你好：" + user.name)
}
