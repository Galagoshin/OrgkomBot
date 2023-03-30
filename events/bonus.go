package events

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func OnBonus(args ...any) {
	user := args[0].(*api.User)
	user_chat := chats.UserChat(user.VKUser)
	bonus1 := user.GetAchievement(api.BONUS_1)
	bonus2 := user.GetAchievement(api.BONUS_2)
	bonus3 := user.GetAchievement(api.BONUS_3)
	if bonus1.GetProgress() < bonus1.GetLimit() {
		bonus1.AddProgress(1)
	}
	if bonus1.GetProgress() == bonus1.GetLimit() && !bonus1.IsCompleted() {
		reward := bonus1.GetReward()
		name := bonus1.GetName()
		user.AddCoins(reward)
		user.AddRating(10)
		bonus1.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if bonus2.GetProgress() < bonus2.GetLimit() {
		bonus2.AddProgress(1)
	}
	if bonus2.GetProgress() == bonus2.GetLimit() && !bonus2.IsCompleted() {
		reward := bonus2.GetReward()
		name := bonus2.GetName()
		user.AddCoins(reward)
		user.AddRating(20)
		bonus2.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if bonus3.GetProgress() < bonus3.GetLimit() && !bonus3.IsCompleted() {
		bonus3.AddProgress(1)
	}
	if bonus3.GetProgress() == bonus3.GetLimit() {
		reward := bonus3.GetReward()
		name := bonus3.GetName()
		user.AddCoins(reward)
		user.AddRating(30)
		bonus3.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
}
