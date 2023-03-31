package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/templates"
	"orgkombot/api"
	"strconv"
	"strings"
)

func StopScanner(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	api.RemoveScannerAmount(user)
	api.RemoveScannerPosition(user)
	api.RemoveScannerEvent(user)
	AdminMenu(chat, outgoing, user)
}

func Scan(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if user.GetAdminLevel() < 1 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	token := outgoing.Text
	amount, exists := api.GetScannerAmount(user)
	position, exists2 := api.GetScannerPosition(user)
	event, exists3 := api.GetScannerEvent(user)
	user.Write(api.Scanner)
	if !exists {
		chat.SendMessage(chats.Message{
			Text: "Произошла какая-то ошибка при выдаче валюты.",
		})
		return
	}
	if !exists2 {
		chat.SendMessage(chats.Message{
			Text: "Произошла какая-то ошибка при выдаче валюты.",
		})
		return
	}
	if !exists3 {
		chat.SendMessage(chats.Message{
			Text: "Произошла какая-то ошибка при выдаче валюты.",
		})
		return
	}
	if event.IsCompleted() {
		chat.SendMessage(chats.Message{
			Text: "Произошла какая-то ошибка при выдаче валюты.",
		})
		return
	}
	if api.GiveVisitByToken(&user, event, token, amount, position) {
		chat.SendMessage(chats.Message{
			Text: "Валюта успешно выдана.",
		})
	} else {
		chat.SendMessage(chats.Message{
			Text: "Некорректный токен! Возможно, QR был плохо отсканирован, либо уже использовался.",
		})
	}
}

func StartScanner(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	args := strings.Split(outgoing.Text, " ")
	events_list := ""
	all_events := api.GetAllEvents()
	for i, event := range all_events {
		events_list += fmt.Sprintf("%d. %s\n", i, event.Name)
	}
	printUsage := func() {
		chat.SendMessage(chats.Message{
			Text: fmt.Sprintf("Usage: /scanner <amount> <event> <position>\nПример для выдачи 100 валюты за посещение на \"%s\": /scanner 100 1 50\n\nСписок меро:\n%s\n\nПодсказка: если начисление валюты идёт за посещение и только, то <position> должен быть равен кол-ву участников на мероприятии. В остальном <position> - это место, занятое на мероприятии.", all_events[1].Name, events_list),
		})
	}
	if user.GetAdminLevel() < 1 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	if len(args) == 4 {
		amount, err := strconv.Atoi(args[1])
		if err != nil {
			printUsage()
			return
		}
		event_id, err := strconv.Atoi(args[2])
		if err != nil {
			printUsage()
			return
		}
		position, err := strconv.Atoi(args[3])
		if err != nil {
			printUsage()
			return
		}
		if position <= 0 || position >= 150 {
			printUsage()
			return
		}
		if amount <= 0 {
			printUsage()
			return
		}
		found := false
		var event *api.Event
		for _, event_el := range api.GetAllEvents() {
			if event_el.Id == uint8(event_id) {
				event = event_el
				found = true
				break
			}
		}
		if !found {
			printUsage()
			return
		}
		user.Write(api.Scanner)
		api.SetScannerAmount(user, uint(amount))
		api.SetScannerPosition(user, uint(position))
		api.SetScannerEvent(user, event)
		kbrd := keyboards.StaticKeyboard{}
		kbrd.Init()
		kbrd.AddButton(keyboards.NormalButton{
			Row:    0,
			Column: 0,
			Payload: keyboards.Payload{
				"action": "scanner stop",
			},
			Color: keyboards.RedColor,
			Text:  "Прекратить выдачу валюты",
		})
		chat.SendMessage(chats.Message{
			Text:     "Пиши сюда токены из QR кодов участников.",
			Keyboard: &kbrd,
		})
	} else {
		printUsage()
	}
}

func GiveAdmin(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	args := strings.Split(outgoing.Text, " ")
	printUsage := func() {
		chat.SendMessage(chats.Message{
			Text: "Usage: /give_admin <link> <level>",
		})
	}
	if user.GetAdminLevel() < 3 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	if len(args) == 3 {
		linked := api.GetUserByLink(args[1])
		if linked != nil {
			linked.Init()
			lvl, err := strconv.Atoi(args[2])
			if err != nil {
				printUsage()
				return
			}
			linked.SetAdminLevel(uint(lvl))
			chat.SendMessage(chats.Message{
				Text: "Админка успешно выдана.",
			})
		} else {
			chat.SendMessage(chats.Message{
				Text: "Пользователь не найден.",
			})
		}
	} else {
		printUsage()
	}
}

func FinishEvent(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if user.GetAdminLevel() < 2 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	event_id := uint8(outgoing.Payload["event"].(float64))
	found := false
	var event *api.Event
	for _, event_el := range api.GetAllEvents() {
		if event_el.Id == event_id {
			event = event_el
			found = true
			break
		}
	}
	if !found {
		chat.SendMessage(chats.Message{
			Text: "Мероприятие не найдено",
		})
		return
	}
	event.Complete()
	chat.SendMessage(chats.Message{
		Text: "Мероприятие завершено",
	})
}

func FinishEventInfo(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if user.GetAdminLevel() < 2 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	events := api.GetAllEvents()
	index := 0
	for i := 0; i < (len(events)+10-1)/10; i++ {
		tmplt := templates.Carousel{}
		tmplt.Init()
		for j := index; j < len(events); j++ {
			event := events[index]
			kbrds := []keyboards.Button{
				keyboards.NormalButton{
					Row:   0,
					Color: keyboards.RedColor,
					Text:  "Завершить",
					Payload: keyboards.Payload{
						"action": "finish_event",
						"event":  event.Id,
					},
				},
			}
			tmplt.AddElement(index-(i*10), event.Name, event.Time, attachments.Image{}, kbrds)
			index++
			if index%10 == 0 {
				break
			}
		}
		chat.SendMessage(chats.Message{Text: fmt.Sprintf("Мероприятия (стр. %d):", i+1), Template: &tmplt})
	}
}

func AdminMenu(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	if user.GetAdminLevel() == 0 {
		chat.SendMessage(chats.Message{
			Text: "У тебя нет прав на выполнение этой команды.",
		})
		return
	}
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "menu",
		},
		Text: "Назад в меню",
	})
	if user.GetAdminLevel() >= 1 {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    1,
			Column: 0,
			Color:  keyboards.BlueColor,
			Payload: keyboards.Payload{
				"action": "scanner",
			},
			Text: "Выдача валюты",
		})
	}
	if user.GetAdminLevel() >= 2 {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    2,
			Column: 0,
			Color:  keyboards.BlueColor,
			Payload: keyboards.Payload{
				"action": "finish_event",
			},
			Text: "Завершение мероприятия",
		})
	}
	if user.GetAdminLevel() >= 3 {
		kbrd.AddButton(keyboards.NormalButton{
			Row:    3,
			Column: 0,
			Color:  keyboards.BlueColor,
			Payload: keyboards.Payload{
				"action": "give_admin",
			},
			Text: "Выдача админки",
		})
	}
	chat.SendMessage(chats.Message{
		Text:     "Выбери на клавиатуре снизу, что хочешь сделать.",
		Keyboard: &kbrd,
	})
}
