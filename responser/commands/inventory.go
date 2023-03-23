package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/templates"
	"orgkombot/api"
)

func Inventory(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	items := user.GetInventory().GetItems()
	if len(items) == 0 {
		chat.SendMessage(chats.Message{Text: "Твой инвентарь пуст."})
		return
	}
	index := 0
	for i := 0; i < (len(items)+10-1)/10; i++ {
		tmplt := templates.Carousel{}
		tmplt.Init()
		for j := index; j < len(items); j++ {
			item := items[index]
			kbrds := []keyboards.Button{
				keyboards.NormalButton{
					Row:   0,
					Color: keyboards.RedColor,
					Text:  "Продать",
					Payload: keyboards.Payload{
						"action": "inventory sell",
						"item":   item.Id,
					},
				},
			}
			if item.Id == api.STICKER_CASE {
				kbrds = []keyboards.Button{
					keyboards.NormalButton{
						Row:   0,
						Color: keyboards.RedColor,
						Text:  "Продать",
						Payload: keyboards.Payload{
							"action": "inventory sell",
							"item":   item.Id,
						},
					},
					keyboards.NormalButton{
						Row:   0,
						Color: keyboards.BlueColor,
						Text:  "Открыть",
						Payload: keyboards.Payload{
							"action": "inventory sell",
							"item":   item.Id,
						},
					},
				}
			} else if item.Id == api.STICKER_CASE_KEY {
				kbrds = []keyboards.Button{}
			}
			tmplt.AddElement(index-(i*10), item.GetName(), item.GetDescription(), item.GetImage(), kbrds)
			index++
			if index%10 == 0 {
				break
			}
		}
		chat.SendMessage(chats.Message{Text: fmt.Sprintf("Инвентарь (стр. %d):", i+1), Template: &tmplt})
	}
}
