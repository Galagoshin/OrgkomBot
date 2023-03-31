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
	kbrd.AddButton(keyboards.NormalButton{
		Row:    1,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "qr",
		},
		Color: keyboards.BlueColor,
		Text:  "QR код 🔐",
	})
	if user.IsSubscribed() {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    3,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "subscribe",
			},
			Color: keyboards.BlueColor,
			Text:  "Отписаться от рассылки 🔕",
		})
	} else {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    3,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "subscribe",
			},
			Color: keyboards.BlueColor,
			Text:  "Подписаться на рассылку 🔔",
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
			visited := "Ни одно мероприятие не было посещено."
			events := user.GetVisitedEvents()
			if len(events) != 0 {
				visited = "Посещённые мероприятия:\n"
				for event, position := range events {
					weight := event.GetWeight()
					ratings_points := weight * 2 * (2.0 / (2.05 * (float64(position+1) - 1.0)))
					points_str := ""
					rated := "*"
					if event.IsRated() {
						points_str = fmt.Sprintf("+%.2f 🏆", ratings_points)
						rated = ""
					}
					visited += fmt.Sprintf("- %s (%.2f%s) %s\n", event.Name, weight, rated, points_str)
				}
			}
			chat.SendMessage(chats.Message{
				Text:     fmt.Sprintf("Твой рейтинг: %.2f 🏆\nТвои коины: %d \U0001FA99\n\n%s", user.GetRating(), user.GetCoins(), visited),
				Keyboard: &kbrd,
			})
		}
	}
}
