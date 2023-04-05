package commands

import (
	"fmt"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
	events2 "orgkombot/events"
	"strconv"
	"strings"
)

func Bank(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "pay",
		},
		Color: keyboards.BlueColor,
		Text:  "Сделать перевод 💸",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    1,
		Column: 0,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "menu",
		},
		Text: "Назад 🔙",
	})
	chat.SendMessage(chats.Message{
		Text:     "Выбери, что хочешь сделать.",
		Keyboard: &kbrd,
	})
}

func StartPay(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "pay cancel",
		},
		Color: keyboards.RedColor,
		Text:  "Отменить перевод 🚫",
	})
	user.Write(api.TypePayUser)
	chat.SendMessage(chats.Message{
		Text:     "Введи ссылку на пользователя, которому хочешь перевести коины.\n\nНапример: vk.com/galagoshin",
		Keyboard: &kbrd,
	})
}

func ChooseAmount(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	linked := api.GetUserByLink(outgoing.Text)
	if linked != nil {
		user.Write(api.TypePayAmount)
		linked.Init()
		api.SetPayUser(user, linked)
		chat.SendMessage(chats.Message{Text: "Введи, сколько коинов хочешь перевести."})
	} else {
		chat.SendMessage(chats.Message{Text: "Пользователей не найден в базе данных, попробуй ещё раз."})
		user.Write(api.TypePayUser)
	}
}

func FinishPay(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	amount, err := strconv.Atoi(outgoing.Text)
	receiver, exists := api.GetPayUser(user)
	if !exists || receiver.GetId() == user.GetId() {
		chat.SendMessage(chats.Message{Text: "При переводе произошла какая-то ошибка."})
		return
	}
	if err != nil {
		chat.SendMessage(chats.Message{Text: "Неправильный формат. Нужно ввести число. Попробуй ещё раз.\n\nПример: 5"})
		user.Write(api.TypePayAmount)
		return
	}
	if amount <= 0 || amount > 1000000 {
		chat.SendMessage(chats.Message{Text: "Неправильный формат. Нужно ввести число > 0. Попробуй ещё раз.\n\nНапример: 5"})
		user.Write(api.TypePayAmount)
		return
	}
	if !user.HaveCoins(uint(amount)) {
		chat.SendMessage(chats.Message{Text: "У тебя нет столько коинов, укажи число меньше ещё раз."})
		user.Write(api.TypePayAmount)
		return
	}
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "pay confirm",
			"amount": amount,
		},
		Color: keyboards.GreenColor,
		Text:  "Да ✅",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 1,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "bank",
		},
		Text: "Нет ❌",
	})
	names := strings.Split(receiver.GetName(), " ")
	first_name := strings.Replace(strings.ToLower(names[0]), string([]rune(strings.ToLower(names[0]))[:1]), strings.ToUpper(string([]rune(names[0])[:1])), 1)
	last_name := strings.Replace(strings.ToLower(names[1]), string([]rune(strings.ToLower(names[1]))[:1]), strings.ToUpper(string([]rune(names[1])[:1])), 1)
	chat.SendMessage(chats.Message{Text: fmt.Sprintf("Уверен, что хочешь перевести %d \U0001FA99 участнику @id%d(%s %s)?", amount, receiver.VKUser, first_name, last_name), Keyboard: &kbrd})
}

func Pay(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	receiver, exists := api.GetPayUser(user)
	api.RemovePayUser(user)
	if !exists || receiver.GetId() == user.GetId() || outgoing.Payload["amount"] == nil {
		chat.SendMessage(chats.Message{Text: "При переводе произошла какая-то ошибка."})
		return
	}
	amount, cast := outgoing.Payload["amount"].(float64)
	if !cast {
		chat.SendMessage(chats.Message{Text: "При переводе произошла какая-то ошибка."})
		return
	}
	if amount <= 0 || amount > 1000000 {
		chat.SendMessage(chats.Message{Text: "При переводе произошла какая-то ошибка."})
		return
	}
	if !user.HaveCoins(uint(amount)) {
		chat.SendMessage(chats.Message{Text: "При переводе произошла какая-то ошибка."})
		return
	}
	receiver.Init()
	user.ReduceCoins(uint(amount))
	receiver.AddCoins(uint(amount))
	kbrd := keyboards.StaticKeyboard{}
	kbrd.Init()
	kbrd.AddButton(keyboards.NormalButton{
		Row:    0,
		Column: 0,
		Payload: keyboards.Payload{
			"action": "pay",
		},
		Color: keyboards.WhiteColor,
		Text:  "Сделать ещё перевод 💸",
	})
	kbrd.AddButton(keyboards.NormalButton{
		Row:    1,
		Column: 0,
		Color:  keyboards.RedColor,
		Payload: keyboards.Payload{
			"action": "menu",
		},
		Text: "Назад в меню 🔙",
	})
	events.CallAllEvents(events2.PayEvent, user, receiver, amount)
	names1 := strings.Split(receiver.GetName(), " ")
	names2 := strings.Split(user.GetName(), " ")
	first_name1 := strings.Replace(strings.ToLower(names1[0]), string([]rune(strings.ToLower(names1[0]))[:1]), strings.ToUpper(string([]rune(names1[0])[:1])), 1)
	last_name1 := strings.Replace(strings.ToLower(names1[1]), string([]rune(strings.ToLower(names1[1]))[:1]), strings.ToUpper(string([]rune(names1[1])[:1])), 1)
	first_name2 := strings.Replace(strings.ToLower(names2[0]), string([]rune(strings.ToLower(names2[0]))[:1]), strings.ToUpper(string([]rune(names2[0])[:1])), 1)
	last_name2 := strings.Replace(strings.ToLower(names2[1]), string([]rune(strings.ToLower(names2[1]))[:1]), strings.ToUpper(string([]rune(names2[1])[:1])), 1)
	chat.SendMessage(chats.Message{Text: fmt.Sprintf("Переведено %d \U0001FA99 участнику @id%d(%s %s)", uint(amount), receiver.VKUser, first_name1, last_name1), Keyboard: &kbrd})
	chats.UserChat(receiver.VKUser).SendMessage(chats.Message{Text: fmt.Sprintf("Тебе перевёл %d \U0001FA99 участник @id%d(%s %s)", uint(amount), user.VKUser, first_name2, last_name2)})
}
