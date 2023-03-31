package events

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func OnPay(args ...any) {
	user := args[0].(api.User)
	user.Init()
	user_chat := chats.UserChat(user.VKUser)
	bank1 := user.GetAchievement(api.BANK_1)
	bank2 := user.GetAchievement(api.BANK_2)
	bank3 := user.GetAchievement(api.BANK_3)
	if bank1.GetProgress() < bank1.GetLimit() {
		bank1.AddProgress(1)
	}
	if bank1.GetProgress() == bank1.GetLimit() && !bank1.IsCompleted() {
		reward := bank1.GetReward()
		name := bank1.GetName()
		user.AddCoins(reward)
		user.AddRating(10)
		bank1.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if bank2.GetProgress() < bank2.GetLimit() {
		bank2.AddProgress(1)
	}
	if bank2.GetProgress() == bank2.GetLimit() && !bank2.IsCompleted() {
		reward := bank2.GetReward()
		name := bank2.GetName()
		user.AddCoins(reward)
		user.AddRating(20)
		bank2.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if bank3.GetProgress() < bank3.GetLimit() {
		bank3.AddProgress(1)
	}
	if bank3.GetProgress() == bank3.GetLimit() && !bank3.IsCompleted() {
		reward := bank3.GetReward()
		name := bank3.GetName()
		user.AddCoins(reward)
		user.AddRating(30)
		bank3.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
}
