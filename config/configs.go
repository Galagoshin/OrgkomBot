package config

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
	"os"
)

func InitAllConfigs() {
	links.Init(map[string]any{
		"vk_group": "https://vk.com/fctk_dm_2023",
		"vk_me":    "https://vk.me/fctk_dm_2023",
	}, 1)
	if !events.Exists() {
		err := events.Create()
		if err != nil {
			logger.Error(err)
			return
		}
		err = events.WriteString("{\"events\": [0:{\"name\": \"\", \"time\": \"\", \"address\": \"\", \"link\": \"\", \"description\": \"\"}]}")
		if err != nil {
			logger.Error(err)
			return
		}
		err = events.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}
	err := events.Open(os.O_RDWR)
	if err != nil {
		logger.Error(err)
		return
	}
	defer func(events *files.File) {
		err := events.Close()
		if err != nil {
			logger.Error(err)
		}
	}(&events)
	cachedEvents = json.Json(events.ReadString())
	if !achievements.Exists() {
		err := achievements.Create()
		if err != nil {
			logger.Error(err)
			return
		}
		err = achievements.WriteString("{\"achievements\": [{\"limit\": 0, \"reward\": 5, \"description\": \"todo something\", \"name\": \"achievement name\"}]}")
		if err != nil {
			logger.Error(err)
			return
		}
		err = achievements.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}
	err = achievements.Open(os.O_RDWR)
	if err != nil {
		logger.Error(err)
		return
	}
	defer func(achievements *files.File) {
		err := achievements.Close()
		if err != nil {
			logger.Error(err)
		}
	}(&achievements)
	cachedAchievements = json.Json(achievements.ReadString())
}
