package events

import "github.com/Galagoshin/GoUtils/events"

const (
	EventVisitEvent     = events.EventName("EventVisitEvent")
	PayEvent            = events.EventName("PayEvent")
	SellItemEvent       = events.EventName("SellItemEvent")
	BuyItemEvent        = events.EventName("BuyItemEvent")
	BonusEvent          = events.EventName("BonusEvent")
	OpenCaseEvent       = events.EventName("OpenCaseEvent")
	GetAchievementEvent = events.EventName("OpenCaseEvent")
	EventCompleteEvent  = events.EventName("EventCompleteEvent")
)
