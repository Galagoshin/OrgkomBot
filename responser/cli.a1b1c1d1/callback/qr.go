package callback

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/api"
	commands2 "orgkombot/responser/cli.a1b1c1d1/commands"
)

func GenQR(args ...any) {
	callback := args[0].(chats.Callback)
	chat := callback.Chat
	user := api.User{VKUser: users.User(chat.GetId())}
	user.Init()
	payload := callback.Payload
	user.Subscribe(!user.IsSubscribed())
	callback.SendAnswer(chats.CallbackAnswer{Text: "Мы генерируем тебе QR, он появится через 2 секунды."})
	if payload["next"] == "profile" {
		go commands2.QR(chat, chats.OutgoingMessage{}, user, true, true)
	} else {
		go commands2.QR(chat, chats.OutgoingMessage{}, user, false, true)
	}
}
