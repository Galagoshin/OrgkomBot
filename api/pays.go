package api

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"net/url"
	"strings"
)

var pays = make(map[User]*User)

func SetPayUser(user User, receiver *User) {
	pays[user] = receiver
}

func GetPayUser(user User) (*User, bool) {
	name, ex := pays[user]
	return name, ex
}

func RemovePayUser(user User) {
	delete(pays, user)
}

func getUserByDomain(domain string) (*User, bool) {
	values := url.Values{
		"v":            {"5.103"},
		"access_token": {tokens.GetToken()},
		"user_ids":     {domain},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/users.get"),
		Data:   values,
	}
	response, err := request.Send()
	if err != nil {
		logger.Error(err)
	}
	response_json, err := json.Decode(json.Json(response.Text()))
	if err != nil {
		logger.Error(err)
	}
	if response_json["error"] != nil {
		logger.Error(errors.New(response_json["error"].(map[string]any)["error_msg"].(string)))
	}
	if response_json["response"] != nil {
		data := response_json["response"].([]any)
		if len(data) == 0 {
			return nil, false
		}
		user_id := data[0].(map[string]any)["id"].(float64)
		user := &User{VKUser: users.User(user_id)}
		user.Init()
		if user.GetId() != 0 {
			return user, true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

func GetUserByLink(link string) *User {
	if len(strings.Split(link, "@")) > 1 && len(strings.Split(strings.Split(link, "@")[1], "]")) > 1 {
		user, exists := getUserByDomain(strings.Split(strings.Split(link, "@")[1], "]")[0])
		if exists {
			return user
		}
	} else if len(strings.Split(link, "vk.com/id")) > 1 {
		user, exists := getUserByDomain(strings.Split(link, "vk.com/")[1])
		if exists {
			return user
		}
	} else if len(strings.Split(link, "vk.com/")) > 1 {
		user, exists := getUserByDomain(strings.Split(link, "vk.com/")[1])
		if exists {
			return user
		}
	} else {
		user, exists := getUserByDomain(link)
		if exists {
			return user
		}
	}
	return nil
}
