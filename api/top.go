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

func GetTopByRating() map[*User]float64 {
	result := make(map[*User]float64)
	rows, err := db.Instance.Query(context.Background(), "SELECT id, vk, coins, rating, full_name, user_group, qr, admin, is_banned, is_subscribed FROM users ORDER BY rating DESC LIMIT 10;")
	if err != nil {
		logger.Error(err)
		return map[*User]float64{}
	}
	for rows.Next() {
		var rating float64
		var id, vk, coins, admin, banned, subscribed uint
		var name, group, qr string
		err := rows.Scan(&id, &vk, &coins, &rating, &name, &group, &qr, &admin, &banned, &subscribed)
		if err != nil {
			logger.Error(err)
			return nil
		}
		var qrcode attachments.Image
		if qr != "nil" {
			qrarr := strings.Split(qr, "_")
			owner_id, err := strconv.Atoi(strings.Split(qrarr[0], "photo")[1])
			if err != nil {
				logger.Error(err)
				return map[*User]float64{}
			}
			idpic, err := strconv.Atoi(qrarr[1])
			if err != nil {
				logger.Error(err)
				return map[*User]float64{}
			}
			qrcode = attachments.Image{
				OwnerId: owner_id,
				Id:      uint(idpic),
			}
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
			admin:  admin,
		}
		if !user.IsBanned() {
			result[user] = rating
		}
	}
	return result
}
