package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
	"strings"
)

func Top(chat chats.Chat, outgoing chats.OutgoingMessage, usr api.User) {
	top := "Топ 10 участников недели:\n"
	i := 1
	is_user_in_top := false
	for _, user := range api.GetTopByRating() {
		if user.GetId() == usr.GetId() {
			is_user_in_top = true
		}
		names := strings.Split(user.GetName(), " ")
		first_name := strings.Replace(strings.ToLower(names[0]), string([]rune(strings.ToLower(names[0]))[:1]), strings.ToUpper(string([]rune(names[0])[:1])), 1)
		last_name := strings.Replace(strings.ToLower(names[1]), string([]rune(strings.ToLower(names[1]))[:1]), strings.ToUpper(string([]rune(names[1])[:1])), 1)
		top += fmt.Sprintf("%d. @id%d(%s %s) - %.2f 🏆\n", i, user.GetVKUser(), first_name, last_name, user.GetRating())
		i++
	}
	if !is_user_in_top {
		names := strings.Split(usr.GetName(), " ")
		first_name := strings.Replace(strings.ToLower(names[0]), string([]rune(strings.ToLower(names[0]))[:1]), strings.ToUpper(string([]rune(names[0])[:1])), 1)
		last_name := strings.Replace(strings.ToLower(names[1]), string([]rune(strings.ToLower(names[1]))[:1]), strings.ToUpper(string([]rune(names[1])[:1])), 1)
		top += fmt.Sprintf("\n.........\n?. @id%d(%s %s) - %.2f 🏆", usr.GetVKUser(), first_name, last_name, usr.GetRating())
	}
	chat.SendMessage(chats.Message{Text: top})
}
