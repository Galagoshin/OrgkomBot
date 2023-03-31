package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func Shop(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	chat.SendMessage(chats.Message{Text: "Магазин мерча станет доступен совсем скоро ;)"})
}
