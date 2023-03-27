package commands

import (
	"fmt"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/templates"
	"orgkombot/api"
)

func EventMore(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if outgoing.Payload["event"] != nil {
		event_id := uint8(outgoing.Payload["event"].(float64))
		var event *api.Event
		for _, event_el := range api.GetAllEvents() {
			if event_el.Id == event_id {
				event = event_el
				break
			}
		}
		weight := "не определён"
		if event.IsCompleted() {
			weight = fmt.Sprintf("%d", event.Weight)
		}
		chat.SendMessage(chats.Message{Text: fmt.Sprintf("%s\n----------------\n%s\n\nВремя проведения: %s\nМесто проведения: %s\nВес: %s", event.Name, event.Description, event.Time, event.Address, weight)})
	}
}

func EventsList(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	events := api.GetAllEvents()
	index := 0
	for i := 0; i < (len(events)+10-1)/10; i++ {
		tmplt := templates.Carousel{}
		tmplt.Init()
		for j := index; j < len(events); j++ {
			event := events[index]
			kbrds := []keyboards.Button{
				keyboards.NormalButton{
					Row:   0,
					Color: keyboards.GreenColor,
					Text:  "Подробнее",
					Payload: keyboards.Payload{
						"action": "event",
						"event":  event.Id,
					},
				},
				keyboards.CallbackButton{
					Row:  1,
					Text: "Зарегистрироваться",
					Payload: keyboards.Payload{
						"action": "no-register",
					},
				},
			}
			if event.Link != "" {
				kbrds = []keyboards.Button{
					keyboards.NormalButton{
						Row:   0,
						Color: keyboards.GreenColor,
						Text:  "Подробнее",
						Payload: keyboards.Payload{
							"action": "event",
							"event":  event.Id,
						},
					},
					keyboards.LinkButton{
						Row:  1,
						Text: "Зарегистрироваться",
						Link: event.Link,
					},
				}
			}
			if event.IsCompleted() {
				kbrds = []keyboards.Button{
					keyboards.NormalButton{
						Row:   0,
						Color: keyboards.RedColor,
						Text:  "Подробнее",
						Payload: keyboards.Payload{
							"action": "event",
							"event":  event.Id,
						},
					},
					keyboards.CallbackButton{
						Row:  1,
						Text: "Зарегистрироваться",
						Payload: keyboards.Payload{
							"action": "no-register",
						},
					},
				}
				tmplt.AddElement(index-(i*10), event.Name, "Завершено", attachments.Image{}, kbrds)
			} else {
				tmplt.AddElement(index-(i*10), event.Name, event.Time, attachments.Image{}, kbrds)
			}
			index++
			if index%10 == 0 {
				break
			}
		}
		chat.SendMessage(chats.Message{Text: fmt.Sprintf("Мероприятия (стр. %d):", i+1), Template: &tmplt})
	}
}
