package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
	"orgkombot/strings"
)

func Bonus(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	kbrd := keyboards.InlineKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.LinkButton{
		Row:    0,
		Column: 0,
		Link:   strings.VK_GROUP,
		Text:   "Перейти в группу",
	})
	chat.SendMessage(chats.Message{
		Text:     "За активность в группе ты будешь получать валюту!\nЕсли ты успеешь совершить активность в группе за первые 10 минут после выхода поста, ты получишь x10 к бонусу, так что обязательно подпишись на уведомления в группе!\n\nЗа лайк поста: +1 coin\nЗа лайк в первые 10 минут: +10 coins\n\nЗа коммент под постом: +1 coin\nЗа коммент в первые 10 минут: +10 coins\n\nЗа репост: + 2 coins\nЗа репост в первые 10 минут: +20 coins",
		Keyboard: &kbrd,
	})
}
