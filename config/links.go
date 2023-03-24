package config

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/configs"
	"github.com/Galagoshin/GoUtils/requests"
)

var links = configs.Config{Name: "links"}

func GetVKGroup() requests.URL {
	val, ex := links.Get("vk_group")
	if !ex {
		logger.Error(errors.New("\"vk_group\" is not defined in links.gconf"))
		return ""
	} else {
		return requests.URL(val.(string))
	}
}

func GetVKMe() requests.URL {
	val, ex := links.Get("vk_me")
	if !ex {
		logger.Error(errors.New("\"vk_me\" is not defined in links.gconf"))
		return ""
	} else {
		return requests.URL(val.(string))
	}
}
