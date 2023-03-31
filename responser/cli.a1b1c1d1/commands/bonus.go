package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
	"orgkombot/config"
)

func Bonus(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	kbrd := keyboards.InlineKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.LinkButton{
		Row:    0,
		Column: 0,
		Link:   config.GetVKGroup(),
		Text:   "Перейти в группу",
	})
	chat.SendMessage(chats.Message{
		Text:     "За активность в группе ты будешь получать валюту!\nЕсли ты успеешь совершить активность в группе за первые 10 минут после выхода поста, ты получишь x10 к бонусу, так что обязательно подпишись на уведомления в группе!\n * Бонус будет выдан, если пост вышел менее суток назад!\n\nЗа лайк поста: +1 \U0001FA99\nЗа лайк в первые 10 минут: +10 \U0001FA99\n\nЗа коммент под постом: +1 \U0001FA99\nЗа коммент в первые 10 минут: +10 \U0001FA99",
		Keyboard: &kbrd,
	})
}
