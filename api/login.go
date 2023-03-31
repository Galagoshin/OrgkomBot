package api

var names = make(map[uint]string)

func SetLoginName(user User, name string) {
	names[user.id] = name
}

func GetLoginName(user User) (string, bool) {
	name, ex := names[user.id]
	return name, ex
}

func RemoveLoginName(user User) {
	delete(names, user.id)
}
