package main

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/time"
	"github.com/Galagoshin/VKGoBot/bot"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
	"github.com/Galagoshin/VKGoBot/bot/vk"
	"orgkombot/config"
	"orgkombot/db"
	events3 "orgkombot/events"
	"orgkombot/responser"
	"orgkombot/responser/callback"
	"orgkombot/tasks"
)

const VERSION = "1.0.0-beta1"

func main() {
	logger.Print(fmt.Sprintf("OrgkomBot v%s has been loaded (%f s.)", VERSION, time.MeasureExecution(func() {
		bot.Init()
		db.Init()
		file := files.File{Path: "./qr_codes"}
		if !file.Exists() {
			directory := &files.Directory{Path: "./qr_codes"}
			err := directory.Create()
			if err != nil {
				logger.Error(err)
				return
			}
		}
		config.InitAllConfigs()
		tasks.AutoStatusTask.Run()
		tasks.WidgetTask.Run()
		tasks.EventsCheckerTask.Run()
		events.RegisterEvent(events.Event{Name: events2.MessageCallbackEvent, Execute: callback.Routing})
		events.RegisterEvent(events.Event{Name: events2.HotReloadEvent, Execute: events3.OnHotReload})
		events.RegisterEvent(events.Event{Name: events2.AddLikeEvent, Execute: events3.OnLike})
		events.RegisterEvent(events.Event{Name: events2.AddCommentEvent, Execute: events3.OnComment})
		events.RegisterEvent(events.Event{Name: events3.EventVisitEvent, Execute: events3.OnEventVisit})
		events.RegisterEvent(events.Event{Name: events3.EventCompleteEvent, Execute: events3.OnEventComplete})
		events.RegisterEvent(events.Event{Name: events3.PayEvent, Execute: events3.OnPay})
		events.RegisterEvent(events.Event{Name: events3.BonusEvent, Execute: events3.OnBonus})
		events.RegisterEvent(events.Event{Name: events3.GetAchievementEvent, Execute: events3.OnGettingAchievement})
		vk.GetHandler().RegisterResponser(responser.Responser)
	})))
	bot.Run()
}
