package responser

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
	"orgkombot/responser/commands"
	"strings"
)

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
	user := api.User{VKUser: message.User}
	user.Init()
	if !message.Client.Keyboard || !message.Client.InlineKeyboard || !message.Client.Carousel {
		commands.Unsupported(chat, message, user)
		return
	}
	canGetCallback := func(buttons []string) bool {
		for _, button_action := range message.Client.ButtonActions {
			if button_action == "callback" {
				return true
			}
		}
		return false
	}
	if !canGetCallback(message.Client.ButtonActions) {
		commands.Unsupported(chat, message, user)
		return
	}
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
			commands.Inventory(chat, message, user)
		} else if message.Payload["action"] == "events" {
			commands.EventsList(chat, message, user)
		} else if message.Payload["action"] == "event" {
			commands.EventMore(chat, message, user)
		} else if strings.Split(message.Payload["action"].(string), " ")[0] == "vote" {
			commands.VoteEvent(chat, message, user)
		} else if message.Payload["action"] == "achievements" {
			commands.Achievements(chat, message, user)
		} else if message.Payload["action"] == "bank" {
			commands.Bank(chat, message, user)
		} else if message.Payload["action"] == "pay" {
			commands.StartPay(chat, message, user)
		} else if message.Payload["action"] == "pay confirm" {
			commands.Pay(chat, message, user)
		} else {
			commands.Menu(chat, message, user, false)
		}
		return
	}
	commands.Menu(chat, message, user, false)
}
