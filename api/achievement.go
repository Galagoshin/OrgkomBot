package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"orgkombot/config"
	"orgkombot/db"
)

const (
	EVENT_1      = 0
	EVENT_2      = 0x1
	EVENT_3      = 0x2
	BANK_1       = 0x3
	BANK_2       = 0x4
	BANK_3       = 0x5
	MERCH        = 0x6
	MARKET       = 0x7
	BONUS_1      = 0x8
	BONUS_2      = 0x9
	BONUS_3      = 0xA
	CASE         = 0xB
	ACHIEVEMENTS = 0xC
)

type Achievement struct {
	id       uint8
	progress uint8
}

func (achievement *Achievement) GetName() string {
	data, err := json.Decode(config.GetAllAchievementsJson())
	if err != nil {
		logger.Error(err)
		return ""
	}
	return data["achievements"].([]any)[achievement.id].(map[string]any)["name"].(string)
}

func (achievement *Achievement) GetDescription() string {
	data, err := json.Decode(config.GetAllAchievementsJson())
	if err != nil {
		logger.Error(err)
		return ""
	}
	return data["achievements"].([]any)[achievement.id].(map[string]any)["description"].(string)
}

func (achievement *Achievement) GetLimit() uint8 {
	data, err := json.Decode(config.GetAllAchievementsJson())
	if err != nil {
		logger.Error(err)
		return 0
	}
	return uint8(data["achievements"].([]any)[achievement.id].(map[string]any)["limit"].(float64))
}

func (achievement *Achievement) GetReward() uint {
	data, err := json.Decode(config.GetAllAchievementsJson())
	if err != nil {
		logger.Error(err)
		return 0
	}
	return uint(data["achievements"].([]any)[achievement.id].(map[string]any)["reward"].(float64))
}

func (user *User) GetCompletedAchievements() []*Achievement {
	achievements := []*Achievement{}
	rows, err := db.Instance.Query(context.Background(), "SELECT achievement_id FROM achievements WHERE owner_id = $1 AND progress = -1;", user.GetId())
	if err != nil {
		logger.Error(err)
		return []*Achievement{}
	}
	for rows.Next() {
		var a_id uint8
		err := rows.Scan(&a_id)
		if err != nil {
			logger.Error(err)
			return nil
		}
		achievement := &Achievement{id: a_id}
		achievement.progress = achievement.GetLimit()
		achievements = append(achievements, achievement)
	}
	return achievements
}

func (user *User) GetAchievements() []*Achievement {
	achievements := []*Achievement{}
	rows, err := db.Instance.Query(context.Background(), "SELECT achievement_id, progress FROM achievements WHERE owner_id = $1;", user.GetId())
	if err != nil {
		logger.Error(err)
		return []*Achievement{}
	}
	for rows.Next() {
		var a_id, progress uint8
		err := rows.Scan(&a_id, &progress)
		if err != nil {
			logger.Error(err)
			return nil
		}
		achievements = append(achievements, &Achievement{id: a_id, progress: progress})
	}
	return achievements
}

func (achievement *Achievement) SetProgress(progress uint8) {
	err := db.Instance.QueryRow(context.Background(), "UPDATE achievement_id SET progress = $1;", progress).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	achievement.progress = progress
}

func (achievement *Achievement) GetId() uint8 {
	return achievement.id
}

func (achievement *Achievement) GetProgress() uint8 {
	return achievement.progress
}

func (achievement *Achievement) AddProgress(val uint8) {
	achievement.SetProgress(achievement.GetProgress() + val)
}
