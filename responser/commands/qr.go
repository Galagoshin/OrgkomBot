package commands

import (
	"fmt"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"orgkombot/api"
)

func QR(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User, profile bool, gen bool) {
	qr := user.GetQR()
	if qr.OwnerId == 0 {
		qr_file := files.File{Path: fmt.Sprintf("./qr_codes/%d.png", user.GetId())}
		if !qr_file.Exists() {
			qr_file = user.GenerateQr()
		}
		qr = chat.UploadImage(qr_file)
	}
	user.SetQR(qr)
	chat.SendMessage(chats.Message{
		Text: "Твой QR код:",
		Attachments: []attachments.Attachment{
			&qr,
		},
	})
	if gen {
		if profile {
			Profile(chat, outgoing, user, false, true)
		} else {
			Menu(chat, outgoing, user, true)
		}
	}
}
