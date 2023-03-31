package responser

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
)

func GetClientSignature(client chats.Client) string {
	signature := ""
	if client.Keyboard {
		signature += "a1"
	} else {
		signature += "a0"
	}
	canGetCallback := func(buttons []string) bool {
		for _, button_action := range client.ButtonActions {
			if button_action == "callback" {
				return true
			}
		}
		return false
	}
	canGetText := func(buttons []string) bool {
		for _, button_action := range client.ButtonActions {
			if button_action == "text" {
				return true
			}
		}
		return false
	}
	if canGetCallback(client.ButtonActions) {
		signature += "b1"
	} else {
		signature += "b0"
	}
	if canGetText(client.ButtonActions) {
		signature += "c1"
	} else {
		signature += "c0"
	}
	if client.Carousel {
		signature += "d1"
	} else {
		signature += "d0"
	}
	return signature
}
