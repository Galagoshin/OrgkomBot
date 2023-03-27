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
		if achievement.GetProgress() == achievement.GetLimit() {
			msg += fmt.Sprintf("⭐ %s 100%s (%d/%d)\n - %s\n\n", achievement.GetName(), "%", achievement.GetProgress(), achievement.GetLimit(), achievement.GetDescription())
		} else {
			msg += fmt.Sprintf("🚫 %s %.1f%s (%d/%d)\n - %s \n\n", achievement.GetName(), (float64(achievement.GetProgress())/float64(achievement.GetLimit()))*100, "%", achievement.GetProgress(), achievement.GetLimit(), achievement.GetDescription())
		}
		//TODO: delete here
		if achievement.GetId() == 0xC {
			break
		}
	}
	chat.SendMessage(chats.Message{Text: msg})
}
