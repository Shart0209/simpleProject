package db

import (
	"simpleProject/pkg/model"
	"time"
)

var DataBaseTest = map[int]model.DocumentManagement{
	1: {
		Id:          1,
		Name:        "Оказание услуг по передаче данных по широкополосному каналу Интернет",
		Date:        time.Date(2022, 12, 15, 0, 0, 0, 0, time.Local),
		Description: "Описание",
		File:        []byte{},
		DateIn:      time.Time{},
		DateOut:     time.Time{},
		Distributor: "пао ростелеком",
		Method:      "аукцион",
		Sum:         150000.00,
	},
	2: {
		Id:          2,
		Name:        "Оказание услуг по передаче данных по широкополосному каналу Интернет",
		Date:        time.Date(2022, 12, 28, 0, 0, 0, 0, time.Local),
		Description: "Описание",
		File:        []byte{},
		DateIn:      time.Time{},
		DateOut:     time.Time{},
		Distributor: "пао ростелеком",
		Method:      "аукцион",
		Sum:         1500.00,
	},
}
