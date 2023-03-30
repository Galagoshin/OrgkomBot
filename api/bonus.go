package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"orgkombot/db"
)

func (user *User) CreatePost(post attachments.Post) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO bonus (post_id, user_id, liked, comment) VALUES ($1, $2, 0, 0);", post.Id, user.GetId()).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (user *User) IsPostCreated(post attachments.Post) bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT id FROM bonus WHERE post_id = $1 AND user_id = $2;", post.Id, user.GetId()).Scan(&val) == nil
}

func (user *User) IsLiked(post attachments.Post) bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT id FROM bonus WHERE post_id = $1 AND user_id = $2 AND liked = 1;", post.Id, user.GetId()).Scan(&val) == nil
}

func (user *User) IsCommented(post attachments.Post) bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT id FROM bonus WHERE post_id = $1 AND user_id = $2 AND comment = 1;", post.Id, user.GetId()).Scan(&val) == nil
}

func (user *User) Like(post attachments.Post) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE bonus SET liked = 1 WHERE post_id = $1 AND user_id = $2;", post.Id, user.GetId()).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (user *User) Dislike(post attachments.Post) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE bonus SET liked = 0 WHERE post_id = $1 AND user_id = $2;", post.Id, user.GetId()).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (user *User) Comment(post attachments.Post) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE bonus SET comment = 1 WHERE post_id = $1 AND user_id = $2;", post.Id, user.GetId()).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (user *User) Uncomment(post attachments.Post) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE bonus SET comment = 0 WHERE post_id = $1 AND user_id = $2;", post.Id, user.GetId()).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}
