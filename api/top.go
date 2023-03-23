package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/db"
	"strconv"
	"strings"
)

func GetTopByRating() map[*User]uint {
	result := make(map[*User]uint)
	rows, err := db.Instance.Query(context.Background(), "SELECT id, vk, coins, rating, full_name, user_group, qr, is_admin, is_banned, is_subscribed FROM users ORDER BY rating DESC LIMIT 10;")
	if err != nil {
		logger.Error(err)
		return map[*User]uint{}
	}
	for rows.Next() {
		var id, vk, coins, admin, banned, rating, subscribed uint
		var name, group, qr string
		err := rows.Scan(&id, &vk, &coins, &rating, &name, &group, &qr, &admin, &banned, &subscribed)
		if err != nil {
			logger.Error(err)
			return nil
		}
		qrarr := strings.Split(qr, "_")
		owner_id, err := strconv.Atoi(strings.Split(qrarr[0], "photo")[1])
		if err != nil {
			logger.Error(err)
			return map[*User]uint{}
		}
		idpic, err := strconv.Atoi(qrarr[1])
		if err != nil {
			logger.Error(err)
			return map[*User]uint{}
		}
		qrcode := attachments.Image{
			OwnerId: owner_id,
			Id:      uint(idpic),
		}
		user := &User{
			VKUser: users.User(vk),
			id:     id,
			name:   name,
			group:  group,
			qr:     qrcode,
			coins:  coins,
			rating: rating,
			banned: banned == 1,
			admin:  admin == 1,
		}
		if !user.IsBanned() {
			result[user] = rating
		}
	}
	return result
}
