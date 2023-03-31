package commands

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/templates"
	"orgkombot/api"
	"strconv"
	"strings"
)

var smiles = map[int]string{
	1: "&#49;&#8419;",
	2: "&#50;&#8419;",
	3: "&#51;&#8419;",
	4: "&#52;&#8419;",
	5: "&#53;&#8419;",
}

func VoteEvent(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if outgoing.Payload["votes"] != nil {
		votes := strings.Split(outgoing.Payload["votes"].(string), ",")
		evid, err := strconv.Atoi(strings.Split(outgoing.Payload["action"].(string), " ")[1])
		if err != nil {
			logger.Error(err)
			return
		}
		event_id := uint8(evid)
		var event *api.Event
		for _, event_el := range api.GetAllEvents() {
			if event_el.Id == event_id {
				event = event_el
				break
			}
		}
		if !event.IsCompleted() || !event.IsVoteOpen() {
			chat.SendMessage(chats.Message{Text: "Отзыв об этом мероприятии больше недоступен."})
			return
		}
		if user.IsVoted(event) {
			chat.SendMessage(chats.Message{Text: "Отзыв о мероприятии уже оставлен."})
			return
		}
		ln := -1
		if outgoing.Payload["votes"].(string) != "" {
			ln = len(votes)
		}
		switch ln {
		case -1:
			kbrd := keyboards.StaticKeyboard{}
			kbrd.Init()
			for i := 5; i >= 1; i-- {
				color := keyboards.WhiteColor
				if i >= 4 {
					color = keyboards.GreenColor
				} else if i <= 2 {
					color = keyboards.RedColor
				}
				kbrd.AddButton(keyboards.NormalButton{
					Row:    0,
					Column: i - 1,
					Color:  color,
					Payload: keyboards.Payload{
						"action": fmt.Sprintf("vote %d", event_id),
						"votes":  fmt.Sprintf("%d", i),
					},
					Text: smiles[i],
				})
			}
			chat.SendMessage(chats.Message{Text: "Как оценишь мероприятие?", Keyboard: &kbrd})
			return
		case 1:
			kbrd := keyboards.StaticKeyboard{}
			kbrd.Init()
			for i := 5; i >= 1; i-- {
				color := keyboards.WhiteColor
				if i >= 4 {
					color = keyboards.GreenColor
				} else if i <= 2 {
					color = keyboards.RedColor
				}
				kbrd.AddButton(keyboards.NormalButton{
					Row:    0,
					Column: i - 1,
					Color:  color,
					Payload: keyboards.Payload{
						"action": fmt.Sprintf("vote %d", event_id),
						"votes":  fmt.Sprintf("%s,%d", outgoing.Payload["votes"].(string), i),
					},
					Text: smiles[i],
				})
			}
			chat.SendMessage(chats.Message{Text: "Как тебе организация мероприятия?", Keyboard: &kbrd})
			return
		case 2:
			kbrd := keyboards.StaticKeyboard{}
			kbrd.Init()
			for i := 5; i >= 1; i-- {
				color := keyboards.WhiteColor
				if i >= 4 {
					color = keyboards.GreenColor
				} else if i <= 2 {
					color = keyboards.RedColor
				}
				kbrd.AddButton(keyboards.NormalButton{
					Row:    0,
					Column: i - 1,
					Color:  color,
					Payload: keyboards.Payload{
						"action": fmt.Sprintf("vote %d", event_id),
						"votes":  fmt.Sprintf("%s,%d", outgoing.Payload["votes"].(string), i),
					},
					Text: smiles[i],
				})
			}
			chat.SendMessage(chats.Message{Text: "Придёшь ли ты на это мероприятие в следующем году?", Keyboard: &kbrd})
			return
		case 3:
			general, err1 := strconv.Atoi(votes[0])
			organization, err2 := strconv.Atoi(votes[1])
			conversion, err3 := strconv.Atoi(votes[2])
			if err1 != nil || err2 != nil || err3 != nil {
				return
			}
			user.VoteEvent(event, general, organization, conversion)
			chat.SendMessage(chats.Message{Text: "Спасибо за отзыв!"})
			Menu(chat, outgoing, user, false)
			return
		}
	}
}

func EventMore(chat chats.Chat, outgoing chats.OutgoingMessage, user api.User) {
	if outgoing.Payload["event"] != nil {
		event_id := uint8(outgoing.Payload["event"].(float64))
		found := false
		var event *api.Event
		for _, event_el := range api.GetAllEvents() {
			if event_el.Id == event_id {
				event = event_el
				found = true
				break
			}
		}
		if !found {
			chat.SendMessage(chats.Message{
				Text: "Мероприятие не найдено",
			})
			return
		}
		weight := "не определён"
		if event.IsCompleted() {
			weight = fmt.Sprintf("%.2f", event.GetWeight())
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
