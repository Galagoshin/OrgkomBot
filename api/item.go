package api

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"strconv"
	"strings"
)

type Item struct {
	Inventory *Inventory
	Id        uint8
}

const (
	STICKER_CASE     = 1
	STICKER_CASE_KEY = 2
)

var descriptions = map[uint8]map[string]any{
	1: {
		"name":  "Стикер-бокс",
		"price": -1,
		"desc":  "Бокс, из которого могут выпасть стикеры",
		"image": "",
	},
	2: {
		"name":  "Ключ",
		"price": 99,
		"desc":  "Ключ к стикер-боксу",
		"image": "",
	},
}

func (item *Item) GetName() string {
	return descriptions[item.Id]["name"].(string)
}

func (item *Item) GetMerchPrice() int {
	return descriptions[item.Id]["price"].(int)
}

func (item *Item) GetDescription() string {
	return descriptions[item.Id]["desc"].(string)
}

func (item *Item) GetImage() attachments.Image {
	image := descriptions[item.Id]["image"].(string)
	if image == "" {
		return attachments.Image{}
	}
	image_arr := strings.Split(image, "_")
	owner_id, err := strconv.Atoi(image_arr[0])
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}
	id, err := strconv.Atoi(image_arr[1])
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}
	return attachments.Image{
		OwnerId: owner_id,
		Id:      uint(id),
	}
}
