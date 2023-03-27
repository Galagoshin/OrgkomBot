package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"orgkombot/config"
	"orgkombot/db"
)

type Event struct {
	Id          uint8
	Name        string
	Time        string
	Address     string
	Description string
	Weight      uint8
	Link        requests.URL
	completed   bool
}

func GetAllEvents() []*Event {
	data, err := json.Decode(config.GetAllEventsJson())
	if err != nil {
		logger.Error(err)
		return []*Event{}
	}
	events := data["events"].([]any)
	result := make([]*Event, len(events))
	for event_id, event_entity := range events {
		result[event_id] = &Event{
			Id:          uint8(event_id),
			Name:        event_entity.(map[string]any)["name"].(string),
			Time:        event_entity.(map[string]any)["time"].(string),
			Address:     event_entity.(map[string]any)["address"].(string),
			Description: event_entity.(map[string]any)["description"].(string),
			Weight:      uint8(event_entity.(map[string]any)["weight"].(float64)),
			Link:        requests.URL(event_entity.(map[string]any)["link"].(string)),
			completed:   false,
		}
		rows, err := db.Instance.Query(context.Background(), "SELECT id FROM events WHERE id = $2;", result[event_id])
		rows.Next()
		rows.Close()
		if err != nil {
			result[event_id].completed = false
		}
		if result[event_id].IsCompleted() {
			var weight uint8
			err := db.Instance.QueryRow(context.Background(), "SELECT weight FROM events WHERE id = $1;", result[event_id]).Scan(&weight)
			if err != nil {
				if err.Error() == "no rows in result set" {
					return []*Event{}
				} else {
					logger.Error(err)
					return []*Event{}
				}
			}
			result[event_id].Weight = weight
		}
	}
	return result
}

func (event *Event) IsCompleted() bool {
	return event.completed
}

func (event *Event) Complete() {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO events (id) VALUES ($1);", event.Id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (user *User) IsVoted(event *Event) bool {
	rows, err := db.Instance.Query(context.Background(), "SELECT user_id FROM events_votes WHERE user_id = $1 AND event_id = $2;", user.GetId(), event.Id)
	rows.Next()
	defer rows.Close()
	return err == nil
}

func (user *User) VoteEvent(event *Event, general, organization, conversion int) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO events_votes (user_id, event_id, general, organization, conversion) VALUES ($1, $2, $3, $4, $5);", user.GetId(), event.Id, general, organization, conversion).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}

func (event *Event) SetWeight() {
	var weight uint8
	err := db.Instance.QueryRow(context.Background(), `
	SELECT ((SUM(general) + SUM(organization) + SUM(conversion)) / 3) / COUNT(*) * 10 + $1 FROM events_votes WHERE event_id = $2;
	`, event.Weight, event.Id).Scan(&weight)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return
		} else {
			logger.Error(err)
			return
		}
	}
	event.Weight = weight
	err = db.Instance.QueryRow(context.Background(), "UPDATE events SET weight = $1;", event.Weight).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}
