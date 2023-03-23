package api

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/VKGoBot/bot/framework"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func (user *User) GenerateQr() files.File {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.id,
	})
	secret, ex := framework.Config.Get("payload-secret-key")
	if !ex {
		logger.Error(errors.New("Secret key not found."))
		return files.File{}
	}
	tokenString, err := token.SignedString([]byte(secret.(string)))
	if err != nil {
		logger.Error(err)
		return files.File{}
	}
	logger.Debug(1, false, tokenString)
	qrc, err := qrcode.New(tokenString)
	if err != nil {
		logger.Error(err)
		return files.File{}
	}
	w0, err := standard.New(fmt.Sprintf("./qr_codes/%d.png", user.GetId()),
		standard.WithHalftone("./logo.png"),
		standard.WithQRWidth(128),
	)
	if err != nil {
		logger.Error(err)
		return files.File{}
	}
	err = qrc.Save(w0)
	if err != nil {
		logger.Error(err)
		return files.File{}
	}
	return files.File{Path: fmt.Sprintf("./qr_codes/%d.png", user.GetId())}
}
