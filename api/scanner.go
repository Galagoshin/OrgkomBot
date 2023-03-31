package api

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/sign"
	"github.com/golang-jwt/jwt/v5"
)

var scanners = make(map[uint]uint)

func SetScannerAmount(user User, amount uint) {
	scanners[user.id] = amount
}

func GetScannerAmount(user User) (uint, bool) {
	amount, ex := scanners[user.id]
	return amount, ex
}

func RemoveScannerAmount(user User) {
	delete(scanners, user.id)
}

func GiveCoinsByToken(admin *User, tokenString string, amount uint) bool {
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
		user.AddCoins(amount)
		CreateTransaction(admin, user, amount)
		return true
	} else {
		return false
	}
}
