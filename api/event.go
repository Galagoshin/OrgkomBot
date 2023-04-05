package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"orgkombot/config"
	"orgkombot/db"
	"strconv"
	"strings"
)

type Event struct {
	Id          uint8
	Name        string
	Time        string
	Address     string
	Description string
	Weight      float64
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
			Weight:      event_entity.(map[string]any)["weight"].(float64),
			Link:        requests.URL(event_entity.(map[string]any)["link"].(string)),
			completed:   true,
		}
		var id uint8
		var weight float64
		err := db.Instance.QueryRow(context.Background(), "SELECT id, weight FROM events WHERE id = $1;", result[event_id].Id).Scan(&id, &weight)
		if err != nil {
			result[event_id].completed = false
			continue
		}
		if result[event_id].IsCompleted() && !result[event_id].IsVoteOpen() {
			result[event_id].Weight = event_entity.(map[string]any)["weight"].(float64)
		}
	}
	return result
}

func (event *Event) GetAllUsers() map[*User]uint {
	result := make(map[*User]uint)
	rows, err := db.Instance.Query(context.Background(), " SELECT users.id, vk, coins, rating, full_name, user_group, qr, admin, is_banned, is_subscribed, events_visits.event_id FROM users LEFT JOIN events_visits ON users.id = user_id WHERE event_id = $1;", event.Id)
	if err != nil {
		logger.Error(err)
		return map[*User]uint{}
	}
	for rows.Next() {
		var rating float64
		var id, vk, coins, admin, banned, subscribed, event_id uint
		var name, group, qr string
		err := rows.Scan(&id, &vk, &coins, &rating, &name, &group, &qr, &admin, &banned, &subscribed, &event_id)
		if err != nil {
			logger.Error(err)
			return nil
		}
		var qrcode attachments.Image
		if qr != "nil" {
			qrarr := strings.Split(qr, "_")
			owner_id, err := strconv.Atoi(strings.Split(qrarr[0], "photo")[1])
			if err != nil {
				logger.Error(err)
				return map[*User]uint{}
			}
			idpic, err := strconv.Atoi(qrarr[1])
			if err != nil {
				logger.Error(err)
				return map[*User]uint{}
			}
			qrcode = attachments.Image{
				OwnerId: owner_id,
				Id:      uint(idpic),
			}
		}
		user := &User{
			VKUser: users.User(vk),
			id:     id,
			name:   name,
			group:  group,
			qr:     qrcode,
			coins:  coins,
			rating: rating,
			banned: banned == 1,
			admin:  admin,
		}
		if !user.IsBanned() {
			result[user] = id
		}
	}
	return result
}

func (event *Event) GetWeight() float64 {
	var weight float64
	err := db.Instance.QueryRow(context.Background(), "SELECT weight FROM events WHERE id = $1;", event.Id).Scan(&weight)
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
		}
		return 0.0
	}
	return weight
}

func (event *Event) IsCompleted() bool {
	return event.completed
}

func (event *Event) Complete() {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO events (id, weight) VALUES ($1, $2);", event.Id, event.Weight).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	events.CallAllEvents("EventCompleteEvent", event)
}

func (user *User) IsVoted(event *Event) bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT user_id FROM events_votes WHERE user_id = $1 AND event_id = $2;", user.GetId(), event.Id).Scan(&val) == nil
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

func (event *Event) IsVoteOpen() bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT id FROM events WHERE created_at > now() - interval '2 hour' AND id = $1;", event.Id).Scan(&val) == nil
}

func (event *Event) IsRated() bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT event_id FROM events_ratings WHERE event_id = $1;", event.Id).Scan(&val) == nil
}

func (event *Event) Rate() {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO events_ratings (event_id) VALUES ($1);", event.Id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	weight := event.GetWeight()
	rows, err := db.Instance.Query(context.Background(), "SELECT id, event_id, user_id, position FROM events_visits WHERE event_id = $1;", event.Id)
	if err != nil {
		logger.Error(err)
		return
	}
	for rows.Next() {
		var id, event_id, user_id, position uint
		err := rows.Scan(&id, &event_id, &user_id, &position)
		if err != nil {
			logger.Error(err)
			return
		}
		rating := weight * 2 * (2.0 / (2.05 * (float64(position+1) - 1.0)))
		err = db.Instance.QueryRow(context.Background(), "UPDATE users SET rating = rating + $1 WHERE id = $2 AND id IN (SELECT user_id FROM events_visits WHERE event_id = $3)", rating, user_id, event.Id).Scan()
		if err != nil {
			if err.Error() != "no rows in result set" {
				logger.Error(err)
				return
			}
		}
	}
	rows.Close()
}

func (user *User) GetVisitedEvents() map[*Event]uint {
	res := make(map[*Event]uint)
	events := GetAllEvents()
	rows, err := db.Instance.Query(context.Background(), "SELECT id, event_id, user_id, position FROM events_visits WHERE user_id = $1;", user.GetId())
	if err != nil {
		logger.Error(err)
		return map[*Event]uint{}
	}
	for rows.Next() {
		var id, event_id, user_id, position uint
		err := rows.Scan(&id, &event_id, &user_id, &position)
		if err != nil {
			logger.Error(err)
			return map[*Event]uint{}
		}
		res[events[event_id]] = position
	}
	rows.Close()
	return res
}

func (user *User) IsVisited(event *Event) bool {
	var val uint
	return db.Instance.QueryRow(context.Background(), "SELECT event_id FROM events_visits WHERE event_id = $1 AND user_id = $2;", event.Id, user.GetId()).Scan(&val) == nil
}

func (user *User) Visit(event *Event, position uint) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO events_visits (event_id, user_id, position) VALUES ($1, $2, $3);", event.Id, user.GetId(), position).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	events.CallAllEvents("EventVisitEvent", event, user, position)
}

func (event *Event) SetWeight() {
	var weight float64
	err := db.Instance.QueryRow(context.Background(), `
	SELECT ((SUM(general) + SUM(organization) + SUM(conversion)) / 3::REAL) / COUNT(*) * 10 + $1 FROM events_votes WHERE event_id = $2;
	`, event.Weight, event.Id).Scan(&weight)
	if err != nil {
		return
	}
	event.Weight = weight
	err = db.Instance.QueryRow(context.Background(), "UPDATE events SET weight = $1 WHERE id = $2;", event.Weight, event.Id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}
