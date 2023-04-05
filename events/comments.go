package events

import (
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/comments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/api"
	"time"
)

func OnComment(args ...any) {
	comment := args[0].(comments.Comment)
	chat := chats.UserChat(comment.Commentator)
	user := &api.User{VKUser: users.User(comment.Commentator)}
	user.Init()
	if user.GetId() == 0 {
		return
	}
	obj := comment.CommentObject
	post := attachments.Post{
		Id:        obj.GetId(),
		OwnerId:   obj.GetOwnerId(),
		AccessKey: obj.GetAccessKey(),
	}
	if !user.IsPostCreated(post) {
		user.CreatePost(post)
	}
	if !user.IsCommented(post) {
		post.Init()
		if time.Now().Unix()-post.Date <= int64(time.Minute.Seconds()*10) {
			chat.SendMessage(chats.Message{Text: "+ 10 \U0001FA99 за коммент в первые 10 минут."})
			user.AddCoins(10)
			user.Comment(post)
			events.CallAllEvents(BonusEvent, user)
		} else if time.Now().Unix()-post.Date <= int64(time.Hour.Seconds()*24) {
			chat.SendMessage(chats.Message{Text: "+ 1 \U0001FA99 за коммент."})
			user.AddCoins(1)
			user.Comment(post)
			events.CallAllEvents(BonusEvent, user)
		}
	}
}
