package responser

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	cli_a1b0c1d0 "orgkombot/responser/cli.a1b0c1d0"
	cli_a1b1c1d1 "orgkombot/responser/cli.a1b1c1d1"
	cli_default "orgkombot/responser/cli.default"
)

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
	signature := GetClientSignature(message.Client)
	logger.Debug(1, false, fmt.Sprintf("Client signature: %s", signature))
	switch signature {
	case "a1b1c1d1":
		cli_a1b1c1d1.Responser(chat, message)
		return
	case "a1b0c1d0":
		cli_a1b0c1d0.Responser(chat, message)
		return
	default:
		cli_default.Responser(chat)
		return
	}
}
