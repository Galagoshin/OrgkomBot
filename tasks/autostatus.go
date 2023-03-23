package tasks

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/scheduler"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/handler"
	"time"
)

const (
	startTime = int64(1680314400)
	endTime   = int64(1681556400)
)

var AutoStatusTask = &scheduler.RepeatingTask{
	Duration:   time.Second,
	OnComplete: AutoStatusExecutor,
}

func AutoStatusExecutor(args ...any) {
	time_now := time.Now().Unix()
	if time_now < startTime {
		logger.Debug(0, false, fmt.Sprintf("Status: %s", formatTime(startTime-time_now)))
		//handler.Group.SetStatus(status)
	} else if time_now < endTime {
		logger.Debug(0, false, fmt.Sprintf("Status: %s", formatTime(endTime-time_now)))
	} else {
		handler.Group.SetStatus("")
	}
}

func formatTime(distance int64) string {
	duration := time.Duration(distance) * time.Second
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%d д. %d ч. %d м.", days, hours, minutes)
}
