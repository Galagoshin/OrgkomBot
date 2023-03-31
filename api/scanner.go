package api

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/sign"
	"github.com/golang-jwt/jwt/v5"
)

var scanners_amount = make(map[uint]uint)

func SetScannerAmount(user User, amount uint) {
	scanners_amount[user.id] = amount
}

func GetScannerAmount(user User) (uint, bool) {
	amount, ex := scanners_amount[user.id]
	return amount, ex
}

func RemoveScannerAmount(user User) {
	delete(scanners_amount, user.id)
}

var scanners_position = make(map[uint]uint)

func SetScannerPosition(user User, position uint) {
	scanners_position[user.id] = position
}

func GetScannerPosition(user User) (uint, bool) {
	position, ex := scanners_position[user.id]
	return position, ex
}

func RemoveScannerPosition(user User) {
	delete(scanners_position, user.id)
}

var scanners_event = make(map[uint]*Event)

func SetScannerEvent(user User, event *Event) {
	scanners_event[user.id] = event
}

func GetScannerEvent(user User) (*Event, bool) {
	event, ex := scanners_event[user.id]
	return event, ex
}

func RemoveScannerEvent(user User) {
	delete(scanners_event, user.id)
}

func GiveVisitByToken(admin *User, event *Event, tokenString string, amount, position uint) bool {
	type UserQR struct {
		Id uint `json:"id"`
		jwt.RegisteredClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserQR{}, func(token *jwt.Token) (any, error) {
		return []byte(sign.SECRET), nil
	})
	if err != nil {
		return false
	}
	if user_token_strct, ok := token.Claims.(*UserQR); ok && token.Valid {
		user := &User{id: user_token_strct.Id}
		user.InitById()
		if user.IsVisited(event) {
			user.AddCoins(amount)
			user.Visit(event, position)
			CreateTransaction(admin, user, amount)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
