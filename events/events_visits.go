package events

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func OnEventVisit(args ...any) {
	user := args[1].(*api.User)
	user_chat := chats.UserChat(user.VKUser)
	event1 := user.GetAchievement(api.EVENT_1)
	event2 := user.GetAchievement(api.EVENT_2)
	event3 := user.GetAchievement(api.EVENT_3)
	if event1.GetProgress() < event1.GetLimit() {
		event1.AddProgress(1)
	}
	if event1.GetProgress() == event1.GetLimit() && !event1.IsCompleted() {
		reward := event1.GetReward()
		name := event1.GetName()
		user.AddCoins(reward)
		user.AddRating(10)
		event1.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if event2.GetProgress() < event2.GetLimit() {
		event2.AddProgress(1)
	}
	if event2.GetProgress() == event2.GetLimit() && !event2.IsCompleted() {
		reward := event2.GetReward()
		name := event2.GetName()
		user.AddCoins(reward)
		user.AddRating(20)
		event2.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
	if event3.GetProgress() < event3.GetLimit() {
		event3.AddProgress(1)
	}
	if event3.GetProgress() == event3.GetLimit() && !event3.IsCompleted() {
		reward := event3.GetReward()
		name := event3.GetName()
		user.AddCoins(reward)
		user.AddRating(30)
		event3.Complete()
		user_chat.SendMessage(chats.Message{Text: fmt.Sprintf("Получено достижение \"%s\" +%d \U0001FA99", name, reward)})
	}
}
