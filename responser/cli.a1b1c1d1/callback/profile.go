package callback

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/api"
	commands2 "orgkombot/responser/cli.a1b1c1d1/commands"
)

func Subscribe(args ...any) {
	callback := args[0].(chats.Callback)
	chat := callback.Chat
	user := api.User{VKUser: users.User(chat.GetId())}
	user.Init()
	user.Subscribe(!user.IsSubscribed())
	commands2.Profile(chat, chats.OutgoingMessage{}, user, true, false)
	callback.SendAnswer(chats.CallbackAnswer{Text: "Изменения профиля сохраненены."})
}
