package cli_default

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
)

func Responser(chat chats.Chat) {
	chat.SendMessage(chats.Message{Text: "Бот не поддерживает твой клиент ВКонтакте! Если ты видишь это сообщение, значит тебе надо обновить ВКонтакте или зайти через браузер."})
}
