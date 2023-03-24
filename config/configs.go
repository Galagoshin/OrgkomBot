package config

import "github.com/Galagoshin/GoLogger/logger"

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
}
