package events

import (
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/likes"
	"orgkombot/api"
	"time"
)

func OnLike(args ...any) {
	like := args[0].(likes.Like)
	chat := chats.UserChat(like.Liker)
	user := &api.User{VKUser: like.Liker}
	user.Init()
	obj := like.LikedObject
	post := attachments.Post{
		Id:        obj.GetId(),
		OwnerId:   obj.GetOwnerId(),
		AccessKey: obj.GetAccessKey(),
	}
	if !user.IsPostCreated(post) {
		user.CreatePost(post)
	}
	if !user.IsLiked(post) {
		post.Init()
		if time.Now().Unix()-post.Date <= int64(time.Minute.Seconds()*10) {
			chat.SendMessage(chats.Message{Text: "+ 10 \U0001FA99 за лайк в первые 10 минут."})
			user.AddCoins(10)
			user.Like(post)
			events.CallAllEvents(BonusEvent, user)
		} else if time.Now().Unix()-post.Date <= int64(time.Hour.Seconds()*24) {
			chat.SendMessage(chats.Message{Text: "+ 1 \U0001FA99 за лайк."})
			user.AddCoins(1)
			user.Like(post)
			events.CallAllEvents(BonusEvent, user)
		}
	}
}
