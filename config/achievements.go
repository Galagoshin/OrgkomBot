package config

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
	"os"
)

var achievements = files.File{
	Path: "./achievements.gconf",
}

func GetAllAchievementsJson() json.Json {
	err := achievements.Open(os.O_RDWR)
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer func(achievements *files.File) {
		err := achievements.Close()
		if err != nil {
			logger.Error(err)
		}
	}(&achievements)
	return json.Json(achievements.ReadString())
}
