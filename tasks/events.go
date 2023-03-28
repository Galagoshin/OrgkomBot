package tasks

import (
	"github.com/Galagoshin/GoUtils/scheduler"
	"orgkombot/api"
	"time"
)

var EventsCheckerTask = &scheduler.RepeatingTask{
	Duration:   time.Minute,
	OnComplete: EventsChecker,
}

func EventsChecker(args ...any) {
	for _, event := range api.GetAllEvents() {
		if event.IsCompleted() && event.IsVoteOpen() {
			event.SetWeight()
		} else if event.IsCompleted() && !event.IsVoteOpen() && !event.IsRated() {
			event.Rate()
		}
	}
}
