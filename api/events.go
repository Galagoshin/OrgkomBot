package api

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"orgkombot/config"
)

type Event struct {
	Name        string
	Time        string
	Address     string
	Description string
	Link        requests.URL
}

func GetAllEvents() []Event {
	data, err := json.Decode(config.GetAllEventsJson())
	if err != nil {
		logger.Error(err)
		return []Event{}
	}
	events := data["events"].([]any)
	result := make([]Event, len(events))
	for event_id, event_entity := range events {
		result[event_id] = Event{
			Name:        event_entity.(map[string]any)["name"].(string),
			Time:        event_entity.(map[string]any)["time"].(string),
			Address:     event_entity.(map[string]any)["address"].(string),
			Description: event_entity.(map[string]any)["description"].(string),
			Link:        requests.URL(event_entity.(map[string]any)["link"].(string)),
		}
	}
	return result
}
