package events

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func OnGettingAchievement(args ...any) {
	achievement := args[0].(*api.Achievement)
	user := achievement.GetOwner()
	user_chat := chats.UserChat(user.VKUser)
	achievements := user.GetAchievement(api.ACHIEVEMENTS)
	if achievements.GetProgress() < achievements.GetLimit() {
		achievements.AddProgress(1)
	}
	if achievements.GetProgress() == achievements.GetLimit() && !achievements.IsCompleted() {
		reward := achievements.GetReward()
		name := achievements.GetName()
		user.AddCoins(reward)
		user.AddRating(200)
		achievements.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
}
