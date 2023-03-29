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

const VERSION = "1.0.0-alpha5"

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
		vk.GetHandler().RegisterResponser(responser.Responser)
	})))
	bot.Run()
}
