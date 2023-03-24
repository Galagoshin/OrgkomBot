package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func Achievements(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	data := user.GetAchievements()
	msg := "Твои достижения: \n"
	for _, achievement := range data {
		progress := achievement.GetProgress()
		if progress == -1 {
			msg += fmt.Sprintf("⭐ %s %d%s\n - %s\n\n", achievement.GetName(), achievement.GetProgress(), "%", achievement.GetDescription())
		} else {
			msg += fmt.Sprintf("🚫 %s %d%s (+%d \U0001FA99)\n - %s \n\n", achievement.GetName(), achievement.GetProgress(), "%", achievement.GetReward(), achievement.GetDescription())
		}
		break
	}
	chat.SendMessage(chats.Message{Text: msg})
}
