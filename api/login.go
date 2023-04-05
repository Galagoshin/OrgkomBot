package api

import "github.com/Galagoshin/GoLogger/logger"

var names = make(map[uint]string)

func SetLoginName(user User, name string) {
	names[uint(user.VKUser)] = name
}

func GetLoginName(user User) (string, bool) {
	logger.Debug(0, false, uint(user.VKUser))
	name, ex := names[uint(user.VKUser)]
	return name, ex
}

func RemoveLoginName(user User) {
	delete(names, uint(user.VKUser))
}
