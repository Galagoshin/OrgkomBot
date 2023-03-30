package api

import "github.com/Galagoshin/VKGoBot/bot/vk/api/users"

const (
	TypeName      = 0x01
	TypeGroup     = 0x02
	TypePayUser   = 0x03
	TypePayAmount = 0x04
)

var inputs = make(map[users.User]int)

func (user User) Read() int {
	val, ex := inputs[user.VKUser]
	delete(inputs, user.VKUser)
	if ex {
		return val
	} else {
		return 0
	}
}

func (user User) Write(in int) {
	inputs[user.VKUser] = in
}
