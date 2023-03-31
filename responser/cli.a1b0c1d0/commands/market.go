package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func Market(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	chat.SendMessage(chats.Message{Text: "Торговая площадка станет доступна совсем скоро ;)"})
}
