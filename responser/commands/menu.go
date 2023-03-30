package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
)

func Menu(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User, gen bool) {
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "profile",
		},
		Color: keyboards.GreenColor,
		Text:  "Профиль 👤",
	})
	qr := user.GetQR()
	if qr.OwnerId == 0 {
		kbrd.AddButton(keyboards.CallbackButton{
			Row:    0,
			Column: 1,
			Payload: keyboards.Payload{
				"action": "qr",
				"next":   "menu",
			},
			Text: "QR код 🔐",
		})
	} else {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    0,
			Column: 1,
			Payload: keyboards.Payload{
				"action": "qr",
			},
			Color: keyboards.GreenColor,
			Text:  "QR код 🔐",
		})
	}
	kbrd.AddButton(keyboards.NormalButton{
		Row:    1,
		Column: 0,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "events",
		},
		Text: "Мероприятия 🎉",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    2,
		Column: 0,
		Color:  keyboards.BlueColor,
		Payload: keyboards.Payload{
			"action": "bank",
		},
		Text: "Банк 🏛",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    2,
		Column: 1,
		Color:  keyboards.BlueColor,
		Payload: keyboards.Payload{
			"action": "bonus",
		},
		Text: "Бонус 🎁",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    2,
		Column: 2,
		Color:  keyboards.BlueColor,
		Payload: keyboards.Payload{
			"action": "market",
		},
		Text: "Рынок 💹",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    3,
		Column: 1,
		Payload: keyboards.Payload{
			"action": "shop",
		},
		Color: keyboards.BlueColor,
		Text:  "Мерч-шоп 🛒",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    3,
		Column: 0,
		Color:  keyboards.BlueColor,
		Payload: keyboards.Payload{
			"action": "top",
		},
		Text: "Рейтинг 🏆",
	})
	if gen {
		chat.SendMessage(chats.Message{
			Text:     "Теперь твой QR сохранён в базе и будет появляться моментально.",
			Keyboard: &kbrd,
		})
	} else {
		chat.SendMessage(chats.Message{
			Text:     "Выбери на клавиатуре снизу, что хочешь глянуть.",
			Keyboard: &kbrd,
		})
	}
}
