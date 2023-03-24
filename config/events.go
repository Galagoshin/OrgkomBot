package config

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
	"os"
)

var events = files.File{
	Path: "./events.gconf",
}

func GetAllEventsJson() json.Json {
	err := events.Open(os.O_RDWR)
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer func(events *files.File) {
		err := events.Close()
		if err != nil {
			logger.Error(err)
		}
	}(&events)
	return json.Json(events.ReadString())
}
