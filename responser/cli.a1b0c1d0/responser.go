package cli_a1b0c1d0

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
	"orgkombot/responser/cli.a1b0c1d0/commands"
	"strings"
)

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
	user := api.User{VKUser: message.User}
	user.Init()
	switch user.Read() {
	case api.TypeName:
		commands.InputName(chat, message, user)
		return
	case api.TypeGroup:
		commands.InputGroup(chat, message, user)
		return
	case api.TypePayUser:
		if message.Payload["action"] != nil && message.Payload["action"] == "pay cancel" {
			commands.Bank(chat, message, user)
		} else {
			commands.ChooseAmount(chat, message, user)
		}
		return
	case api.TypePayAmount:
		if message.Payload["action"] != nil && message.Payload["action"] == "pay cancel" {
			commands.Bank(chat, message, user)
		} else {
			commands.FinishPay(chat, message, user)
		}
		return
	case api.Scanner:
		if message.Payload["action"] != nil && message.Payload["action"] == "scanner stop" {
			commands.Unsupported(chat, message, user)
		} else {
			commands.Unsupported(chat, message, user)
		}
		return
	}
	if user.GetId() == 0 {
		commands.StartLogin(chat, message, user)
		return
	}
	if message.Payload["action"] != nil {
		if message.Payload["action"] == "qr" {
			commands.QR(chat, message, user, false, false)
		} else if message.Payload["action"] == "qrp" {
			commands.QR(chat, message, user, true, false)
		} else if message.Payload["action"] == "bonus" {
			commands.Bonus(chat, message, user)
		} else if message.Payload["action"] == "cases" {
			commands.Cases(chat, message, user)
		} else if message.Payload["action"] == "top" {
			commands.Top(chat, message, user)
		} else if message.Payload["action"] == "profile" {
			commands.Profile(chat, message, user, false, false)
		} else if message.Payload["action"] == "inventory" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "events" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "event" {
			commands.Unsupported(chat, message, user)
		} else if strings.Split(message.Payload["action"].(string), " ")[0] == "vote" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "achievements" {
			commands.Achievements(chat, message, user)
		} else if message.Payload["action"] == "bank" {
			commands.Bank(chat, message, user)
		} else if message.Payload["action"] == "pay" {
			commands.StartPay(chat, message, user)
		} else if message.Payload["action"] == "pay confirm" {
			commands.Pay(chat, message, user)
		} else if message.Payload["action"] == "market" {
			commands.Market(chat, message, user)
		} else if message.Payload["action"] == "shop" {
			commands.Shop(chat, message, user)
		} else if message.Payload["action"] == "admin" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "give_admin" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "subscribe" {
			commands.Unsupported(chat, message, user)
		} else if message.Payload["action"] == "finish_event" {
			if message.Payload["event"] != nil {
				commands.Unsupported(chat, message, user)
			} else {
				commands.Unsupported(chat, message, user)
			}
		} else if message.Payload["action"] == "scanner" {
			commands.Unsupported(chat, message, user)
		} else {
			commands.Menu(chat, message, user, false)
		}
	} else {
		if strings.Split(message.Text, " ")[0] == "/give_admin" {
			commands.Unsupported(chat, message, user)
		} else if strings.Split(message.Text, " ")[0] == "/scanner" {
			commands.Unsupported(chat, message, user)
		} else {
			commands.Menu(chat, message, user, false)
		}
	}
}
