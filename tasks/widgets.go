package tasks

import (
	"github.com/Galagoshin/GoUtils/scheduler"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/handler"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/widgets"
	"orgkombot/api"
	"orgkombot/config"
	"time"
)

var WidgetTask = &scheduler.RepeatingTask{
	Duration:   time.Second * 5,
	OnComplete: WidgetUpdater,
}

func WidgetUpdater(args ...any) {
	widget := widgets.WidgetComactList{
		Title:      "Мероприятия",
		TitleUrl:   config.GetVKGroup(),
		FooterText: "Подробнее о мероприятиях",
		FooterUrl:  config.GetVKMe(),
	}
	widget.Init()
	count := 0
	for _, event := range api.GetAllEvents() {
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
		count++
		if count == 4 {
			break
		}
	}
	handler.Group.SetWidget(&widget)
}
