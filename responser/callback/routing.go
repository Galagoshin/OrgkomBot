package callback

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
)

func Routing(args ...any) {
	callback := args[0].(chats.Callback)
	payload := callback.Payload
	if payload["action"] != nil {
		if payload["action"] == "subscribe" {
			Subscribe(args...)
			return
		} else if payload["action"] == "no-register" {
			NoRegister(args...)
			return
		}
	}
	if payload["action"] != nil && payload["next"] != nil {
		if payload["action"] == "qr" {
			GenQR(args...)
			return
		}
	}
}
