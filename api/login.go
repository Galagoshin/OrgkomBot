package api

var names = make(map[User]string)

func SetLoginName(user User, name string) {
	names[user] = name
}

func GetLoginName(user User) (string, bool) {
	name, ex := names[user]
	return name, ex
}

func RemoveLoginName(user User) {
	delete(names, user)
}
