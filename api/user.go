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

type User struct {
	VKUser     users.User
	id         uint
	name       string
	group      string
	qr         attachments.Image
	coins      uint
	rating     float64
	admin      uint
	banned     bool
	subscribed bool
}

func GetAllSubscribedUsers() map[*User]uint {
	result := make(map[*User]uint)
	rows, err := db.Instance.Query(context.Background(), "SELECT id, vk, coins, rating, full_name, user_group, qr, admin, is_banned, is_subscribed FROM users WHERE is_subscribed = 1;")
	if err != nil {
		logger.Error(err)
		return map[*User]uint{}
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
				return map[*User]uint{}
			}
			idpic, err := strconv.Atoi(qrarr[1])
			if err != nil {
				logger.Error(err)
				return map[*User]uint{}
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
			result[user] = id
		}
	}
	return result
}

func (user *User) Delete() {
	err := db.Instance.QueryRow(context.Background(), "DELETE FROM users WHERE vk = $1;", user.VKUser).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.id = 0
	user.coins = 0
	user.rating = 0
	user.name = ""
	user.group = ""
	user.qr = attachments.Image{}
	user.admin = 0
	user.banned = false
	user.subscribed = false
}

func (user *User) Create(name, group string) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO users (vk, full_name, user_group) VALUES ($1, $2, $3);", user.VKUser, name, group).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.Init()
	for i := 0; i <= 0x10; i++ {
		err := db.Instance.QueryRow(context.Background(), "INSERT INTO achievements (owner_id, achievement_id) VALUES ($1, $2);", user.GetId(), i).Scan()
		if err != nil {
			if err.Error() != "no rows in result set" {
				logger.Error(err)
				return
			}
		}
	}
}

func (user *User) Init() {
	var rating float64
	var id, vk, coins, admin, banned, subscribed uint
	var name, group, qr string
	err := db.Instance.QueryRow(context.Background(), "SELECT id, vk, coins, rating, full_name, user_group, qr, admin, is_banned, is_subscribed FROM users WHERE vk = $1;", user.VKUser).Scan(&id, &vk, &coins, &rating, &name, &group, &qr, &admin, &banned, &subscribed)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return
		} else {
			logger.Error(err)
			return
		}
	}
	user.id = id
	user.name = name
	user.group = group
	user.coins = coins
	user.rating = rating
	user.banned = banned == 1
	user.admin = admin
	user.subscribed = subscribed == 1
	var qrcode attachments.Image
	if qr == "nil" {
		user.qr = attachments.Image{}
	} else {
		qrarr := strings.Split(qr, "_")
		owner_id, err := strconv.Atoi(strings.Split(qrarr[0], "photo")[1])
		if err != nil {
			logger.Error(err)
			return
		}
		idpic, err := strconv.Atoi(qrarr[1])
		if err != nil {
			logger.Error(err)
			return
		}
		qrcode = attachments.Image{
			OwnerId: owner_id,
			Id:      uint(idpic),
		}
	}
	user.qr = qrcode
}

func (user *User) InitById() {
	var rating float64
	var id, vk, coins, admin, banned, subscribed uint
	var name, group, qr string
	err := db.Instance.QueryRow(context.Background(), "SELECT id, vk, coins, rating, full_name, user_group, qr, admin, is_banned, is_subscribed FROM users WHERE id = $1;", user.id).Scan(&id, &vk, &coins, &rating, &name, &group, &qr, &admin, &banned, &subscribed)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return
		} else {
			logger.Error(err)
			return
		}
	}
	user.id = id
	user.name = name
	user.group = group
	user.coins = coins
	user.rating = rating
	user.banned = banned == 1
	user.admin = admin
	user.subscribed = subscribed == 1
	var qrcode attachments.Image
	if qr == "nil" {
		user.qr = attachments.Image{}
	} else {
		qrarr := strings.Split(qr, "_")
		owner_id, err := strconv.Atoi(strings.Split(qrarr[0], "photo")[1])
		if err != nil {
			logger.Error(err)
			return
		}
		idpic, err := strconv.Atoi(qrarr[1])
		if err != nil {
			logger.Error(err)
			return
		}
		qrcode = attachments.Image{
			OwnerId: owner_id,
			Id:      uint(idpic),
		}
	}
	user.qr = qrcode
}

func (user *User) Ban(value bool) {
	val := 1
	if !value {
		val = 0
	}
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET is_banned = $1 WHERE id = $2;", val, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.banned = value
}

func (user *User) SetCoins(value uint) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET coins = $1 WHERE id = $2;", value, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.coins = value
}

func (user *User) AddCoins(value uint) {
	user.SetCoins(user.GetCoins() + value)
}

func (user *User) ReduceCoins(value uint) {
	user.SetCoins(user.GetCoins() - value)
}

func (user *User) HaveCoins(value uint) bool {
	return int(user.GetCoins())-int(value) >= 0
}

func (user *User) GetCoins() uint {
	return user.coins
}

func (user *User) SetRating(value float64) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET rating = $1 WHERE id = $2;", value, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.rating = value
}

func (user *User) AddRating(value float64) {
	user.SetRating(user.GetRating() + value)
}

func (user *User) ReduceRating(value float64) {
	user.SetRating(user.GetRating() - value)
}

func (user *User) HaveRating(value uint) bool {
	return int(user.GetCoins())-int(value) >= 0
}

func (user *User) GetRating() float64 {
	return user.rating
}

func (user *User) SetName(value string) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET full_name = $1 WHERE id = $2;", value, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.name = value
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) IsBanned() bool {
	return user.banned
}

func (user *User) GetAdminLevel() uint {
	return user.admin
}

func (user *User) SetAdminLevel(value uint) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET admin = $1 WHERE id = $2;", value, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.admin = value
}

func (user *User) Subscribe(value bool) {
	val := 1
	if !value {
		val = 0
	}
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET is_subscribed = $1 WHERE id = $2;", val, user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.subscribed = value
}

func (user *User) IsSubscribed() bool {
	return user.subscribed
}

func (user *User) SetQR(image attachments.Image) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE users SET qr = $1 WHERE id = $2;", image.BuildString(), user.id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	user.qr = image
}

func (user *User) GetQR() attachments.Image {
	return user.qr
}

func (user *User) GetVKUser() users.User {
	return user.VKUser
}

func (user *User) GetId() uint {
	return user.id
}
