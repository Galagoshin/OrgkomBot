package callback

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/api"
)

func NoRegister(args ...any) {
	callback := args[0].(chats.Callback)
	chat := callback.Chat
	user := api.User{VKUser: users.User(chat.GetId())}
	user.Init()
	callback.SendAnswer(chats.CallbackAnswer{Text: "Регистрация на это мероприятие недоступна или отсутствует."})
}
