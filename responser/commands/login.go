package commands

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
	"regexp"
)

func StartLogin(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	chat.SendMessage(chats.Message{Text: "Привет! Напиши своё полное ФИО в формате \"Фамилия Имя Отчество\" в чат и отправь мне.\nНапример: Иванов Иван Иванович"})
	user.Write(api.TypeName)
}

func InputName(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	match, _ := regexp.MatchString("^([А-Яа-яё\\-]{1,25} [А-Яа-яё\\-]{1,25} [А-Яа-яё\\-]{1,25})$", outgoing.Text)
	if !match {
		chat.SendMessage(chats.Message{Text: "Неправильный формат ФИО!\nФормат: \"Фамилия Имя Отчество\".\nНапример: Иванов Иван Иванович"})
		user.Write(api.TypeName)
		return
	}
	api.SetLoginName(user, outgoing.Text)
	chat.SendMessage(chats.Message{Text: "Теперь введи свою группу."})
	user.Write(api.TypeGroup)
}

func InputGroup(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	login_name, exists_login := api.GetLoginName(user)
	if !exists_login {
		chat.SendMessage(chats.Message{Text: "Произошла какая-то ошибка :(\nВведи своё ФИО ещё раз."})
		user.Write(api.TypeName)
		return
	}
	if len(outgoing.Text) >= 20 {
		chat.SendMessage(chats.Message{Text: "Неправильный формат группы! Введи ещё раз.\nНапример: ММБ-004"})
		user.Write(api.TypeGroup)
		return
	}
	api.SetLoginName(user, outgoing.Text)
	user.Create(login_name, outgoing.Text)
	chat.SendMessage(chats.Message{Text: "Отлично! Теперь тебе доступны основные функции бота."})
	Menu(chat, outgoing, user)
}
