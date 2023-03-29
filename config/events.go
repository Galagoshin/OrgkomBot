package config

import (
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
)

var events = files.File{
	Path: "./events.gconf",
}

var cachedEvents json.Json

func GetAllEventsJson() json.Json {
	return cachedEvents
}
