package events

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"orgkombot/api"
)

func OnEventComplete(args ...any) {
	event := args[0].(*api.Event)
	for user, _ := range event.GetAllUsers() {
		user_chat := chats.UserChat(user.VKUser)
		kbrd := keyboards.InlineKeyboard{}
		kbrd.Init()
		kbrd.AddButton(keyboards.NormalButton{
			Row:    0,
			Column: 0,
			Payload: keyboards.Payload{
				"action": fmt.Sprintf("vote %d", event.Id),
				"votes":  "",
			},
			Color: keyboards.GreenColor,
			Text:  "Оставить отзыв 📃",
		})
		user_chat.SendMessage(chats.Message{
			Text:     fmt.Sprintf("Оставь отзыв о мероприятии \"%s\"", event.Name),
			Keyboard: &kbrd,
		})
	}
}
