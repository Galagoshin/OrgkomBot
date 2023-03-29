package config

import (
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
)

var achievements = files.File{
	Path: "./achievements.gconf",
}

var cachedAchievements json.Json

func GetAllAchievementsJson() json.Json {
	return cachedAchievements
}
