package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"orgkombot/db"
)

const (
	EVENT_1      = 1
	EVENT_2      = 2
	EVENT_3      = 3
	BANK_1       = 4
	BANK_2       = 5
	BANK_3       = 6
	MERCH        = 7
	MARKET       = 8
	BONUS_1      = 9
	BONUS_2      = 10
	BONUS_3      = 11
	CASE         = 12
	ACHIEVEMENTS = 13
)

type Achievement struct {
	id       uint8
	progress uint8
}

func (user *User) GetCompletedAchievements() []*Achievement {
	achievements := []*Achievement{}
	rows, err := db.Instance.Query(context.Background(), "SELECT achievement_id FROM achievements WHERE owner_id = $1 AND progress = 100;", user.GetId())
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
		achievements = append(achievements, &Achievement{id: a_id, progress: 100})
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

func (achievement *Achievement) GetProgress() uint8 {
	return achievement.progress
}

func (achievement *Achievement) AddProgress(val uint8) {
	achievement.SetProgress(achievement.GetProgress() + val)
}
