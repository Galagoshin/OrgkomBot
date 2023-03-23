package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
)

func Profile(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User, edited, gen bool) {
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "inventory",
		},
		Color: keyboards.BlueColor,
		Text:  "Инвентарь 👕",
	})
	qr := user.GetQR()
	if qr.OwnerId == 0 {
		kbrd.AddButton(keyboards.CallbackButton{
			Row:    1,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "qr",
				"next":   "profile",
			},
			Text: "QR код 🔐",
		})
	} else {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    1,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "qrp",
			},
			Color: keyboards.BlueColor,
			Text:  "QR код 🔐",
		})
	}
	if user.IsSubscribed() {
		kbrd.AddButton(keyboards.CallbackButton{
			Row:    3,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "subscribe",
			},
			Text: "Отписаться от рассылки 🔕",
		})
	} else {
		kbrd.AddButton(keyboards.CallbackButton{
			Row:    3,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "subscribe",
			},
			Text: "Подписаться на рассылку 🔔",
		})
	}
	kbrd.AddButton(keyboards.NormalButton{
		Row:    2,
		Column: 0,
		Color:  keyboards.BlueColor,
		Payload: keyboards.Payload{
			"action": "achievements",
		},
		Text: "Достижения ⭐",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    4,
		Column: 0,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "menu",
		},
		Text: "Назад 🔙",
	})
	if edited {
		if !user.IsSubscribed() {
			chat.SendMessage(chats.Message{
				Text:     "Вы отписались от рассылки постов и уведомлений.",
				Keyboard: &kbrd,
			})
		} else {
			chat.SendMessage(chats.Message{
				Text:     "Вы подписались на рассылку постов и уведомлений.",
				Keyboard: &kbrd,
			})
		}
	} else {
		if gen {
			chat.SendMessage(chats.Message{
				Text:     fmt.Sprintf("Теперь твой QR сохранён в базе и будет появляться моментально."),
				Keyboard: &kbrd,
			})
		} else {
			chat.SendMessage(chats.Message{
				Text:     fmt.Sprintf("Твой рейтинг: %d 🏆\nТвои коины: %d \U0001FA99", user.GetRating(), user.GetCoins()),
				Keyboard: &kbrd,
			})
		}
	}
}
