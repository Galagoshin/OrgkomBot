package events

import "orgkombot/config"

func OnHotReload(...any) {
	config.InitAllConfigs()
}
