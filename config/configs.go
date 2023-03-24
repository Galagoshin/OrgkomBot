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
		err = events.WriteString("{\"events\": {0:{\"name\": \"\", \"time\": \"\", \"address\": \"\", \"link\": \"\", \"description\": \"\"}}}")
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
}
