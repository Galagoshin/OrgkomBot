package callback

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/api"
	"orgkombot/responser/commands"
)

func Schedule(args ...any) {
	callback := args[0].(chats.Callback)
	chat := callback.Chat
	user := api.User{VKUser: users.User(chat.GetId())}
	user.Init()
	payload := callback.Payload
	if payload["action"] != nil {
		if payload["action"] == "subscribe" {
			if user.IsSubscribed() {
				user.Subscribe(false)
				commands.Profile(chat, chats.OutgoingMessage{}, user, true)
			} else {
				user.Subscribe(true)
				commands.Profile(chat, chats.OutgoingMessage{}, user, true)
			}
			callback.SendAnswer(chats.CallbackAnswer{Text: "Изменения профиля сохраненены."})
			return
		}
	}
}
