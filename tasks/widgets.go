package tasks

import (
	"fmt"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/GoUtils/scheduler"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/handler"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/widgets"
	"orgkombot/api"
	"orgkombot/config"
	"strings"
	"time"
)

var WidgetTask = &scheduler.RepeatingTask{
	Duration:   time.Minute,
	OnComplete: WidgetUpdater,
}

var executors_list = []func(){
	list_widget,
	list_widget,
	list_widget,
	top_widget,
}

var current_executor = 0

func list_widget() {
	widget := widgets.WidgetComactList{
		Title:      "Мероприятия",
		TitleUrl:   config.GetVKGroup(),
		FooterText: "Подробнее о мероприятиях",
		FooterUrl:  config.GetVKMe(),
	}
	widget.Init()
	count := 0
	for _, event := range api.GetAllEvents() {
		if !event.IsCompleted() {
			if event.Link == "" {
				widget.AddRow(widgets.ListRow{
					Index:       count,
					Title:       event.Name,
					TitleUrl:    config.GetVKMe(),
					Time:        event.Time,
					Address:     event.Address,
					Description: event.Description,
				})
			} else {
				widget.AddRow(widgets.ListRow{
					Index:       count,
					Title:       event.Name,
					TitleUrl:    config.GetVKMe(),
					ButtonText:  "Зарегистрироваться",
					ButtonUrl:   event.Link,
					Time:        event.Time,
					Address:     event.Address,
					Description: event.Description,
				})
			}
			count++
			if count == 3 {
				break
			}
		}
	}
	handler.Group.SetWidget(&widget)
}

func top_widget() {
	widget := widgets.WidgetTable{
		Title:               "Лучшие участники недели",
		TitleUrl:            string(config.GetVKMe()),
		ColumnsDescriptions: []string{"Участник", "Рейтинг"},
	}
	widget.Init()
	count := 0
	for user, rating := range api.GetTopByRating() {
		names := strings.Split(user.GetName(), " ")
		first_name := strings.Replace(strings.ToLower(names[0]), string([]rune(strings.ToLower(names[0]))[:1]), strings.ToUpper(string([]rune(names[0])[:1])), 1)
		last_name := strings.Replace(strings.ToLower(names[1]), string([]rune(strings.ToLower(names[1]))[:1]), strings.ToUpper(string([]rune(names[1])[:1])), 1)
		columns := []widgets.TableColumn{
			{
				Index: 0,
				Text:  fmt.Sprintf("%s %s", first_name, last_name),
				Url:   requests.URL(fmt.Sprintf("https://vk.com/id%d", user.VKUser)),
				Icon:  attachments.Image{OwnerId: int(user.VKUser)},
			},
			{
				Index: 1,
				Text:  fmt.Sprintf("%d", rating),
				Url:   requests.URL(fmt.Sprintf("https://vk.com/id%d", user.VKUser)),
			},
		}
		widget.AddRow(widgets.TableRow{
			Index:   count,
			Columns: columns,
		})
	}
	handler.Group.SetWidget(&widget)
}

func WidgetUpdater(...any) {
	executors_list[current_executor]()
	current_executor += 1
	if current_executor == len(executors_list) {
		current_executor = 0
	}
}
