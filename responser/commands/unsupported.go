package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func Unsupported(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	chat.SendMessage(chats.Message{Text: "Бот работает только с официальными клиентами ВКонтакте актуальных версий! Если ты видишь это сообщение, значит тебе надо обновить ВКонтакте или зайти через браузер."})
}
