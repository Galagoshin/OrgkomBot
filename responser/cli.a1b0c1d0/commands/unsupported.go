package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func Unsupported(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	chat.SendMessage(chats.Message{Text: "Эта функция не поддерживается твоим клиентом ВКонтакте! Тебе нужно обновить приложение или зайти через браузер."})
}
