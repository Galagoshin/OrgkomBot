package cli_a1b1c1d1

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
	commands2 "orgkombot/responser/cli.a1b1c1d1/commands"
	"strings"
)

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
	user := api.User{VKUser: message.User}
	user.Init()
	switch user.Read() {
	case api.TypeName:
		commands2.InputName(chat, message, user)
		return
	case api.TypeGroup:
		commands2.InputGroup(chat, message, user)
		return
	case api.TypePayUser:
		if message.Payload["action"] != nil && message.Payload["action"] == "pay cancel" {
			commands2.Bank(chat, message, user)
		} else {
			commands2.ChooseAmount(chat, message, user)
		}
		return
	case api.TypePayAmount:
		if message.Payload["action"] != nil && message.Payload["action"] == "pay cancel" {
			commands2.Bank(chat, message, user)
		} else {
			commands2.FinishPay(chat, message, user)
		}
		return
	case api.Scanner:
		if message.Payload["action"] != nil && message.Payload["action"] == "scanner stop" {
			commands2.StopScanner(chat, message, user)
		} else {
			commands2.Scan(chat, message, user)
		}
		return
	}
	if user.GetId() == 0 {
		commands2.StartLogin(chat, message, user)
		return
	}
	if message.Payload["action"] != nil {
		if message.Payload["action"] == "qr" {
			commands2.QR(chat, message, user, false, false)
		} else if message.Payload["action"] == "qrp" {
			commands2.QR(chat, message, user, true, false)
		} else if message.Payload["action"] == "bonus" {
			commands2.Bonus(chat, message, user)
		} else if message.Payload["action"] == "cases" {
			commands2.Cases(chat, message, user)
		} else if message.Payload["action"] == "top" {
			commands2.Top(chat, message, user)
		} else if message.Payload["action"] == "profile" {
			commands2.Profile(chat, message, user, false, false)
		} else if message.Payload["action"] == "inventory" {
			commands2.Inventory(chat, message, user)
		} else if message.Payload["action"] == "events" {
			commands2.EventsList(chat, message, user)
		} else if message.Payload["action"] == "event" {
			commands2.EventMore(chat, message, user)
		} else if strings.Split(message.Payload["action"].(string), " ")[0] == "vote" {
			commands2.VoteEvent(chat, message, user)
		} else if message.Payload["action"] == "achievements" {
			commands2.Achievements(chat, message, user)
		} else if message.Payload["action"] == "bank" {
			commands2.Bank(chat, message, user)
		} else if message.Payload["action"] == "pay" {
			commands2.StartPay(chat, message, user)
		} else if message.Payload["action"] == "pay confirm" {
			commands2.Pay(chat, message, user)
		} else if message.Payload["action"] == "market" {
			commands2.Market(chat, message, user)
		} else if message.Payload["action"] == "shop" {
			commands2.Shop(chat, message, user)
		} else if message.Payload["action"] == "admin" {
			commands2.AdminMenu(chat, message, user)
		} else if message.Payload["action"] == "give_admin" {
			commands2.GiveAdmin(chat, message, user)
		} else if message.Payload["action"] == "finish_event" {
			if message.Payload["event"] != nil {
				commands2.FinishEvent(chat, message, user)
			} else {
				commands2.FinishEventInfo(chat, message, user)
			}
		} else if message.Payload["action"] == "scanner" {
			commands2.StartScanner(chat, message, user)
		} else {
			commands2.Menu(chat, message, user, false)
		}
	} else {
		if strings.Split(message.Text, " ")[0] == "/give_admin" {
			commands2.GiveAdmin(chat, message, user)
		} else if strings.Split(message.Text, " ")[0] == "/scanner" {
			commands2.StartScanner(chat, message, user)
		} else {
			commands2.Menu(chat, message, user, false)
		}
	}
}
