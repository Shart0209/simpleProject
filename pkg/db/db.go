package db

import (
	"simpleProject/pkg/model"
)

var Data = map[int]model.DocumentManagment{
	1: {
		Id:          "1",
		Name:        "Оказание услуг по передаче данных по широкополосному каналу Интернет",
		Date:        "28.12.2022",
		Description: "Описание",
		File:        "generic UUID file",
		DateIn:      "01.01.2023",
		DateOut:     "31.12.2023",
		Distributor: "пао ростелеком",
		Metod:       "аукцион",
		Sum:         "150 000",
	},
	2: {
		Id:          "2",
		Name:        "Оказание услуг по передаче данных по широкополосному каналу Интернет",
		Date:        "28.12.2022",
		Description: "Описание",
		File:        "generic UUID file",
		DateIn:      "01.01.2023",
		DateOut:     "31.12.2023",
		Distributor: "пао ростелеком",
		Metod:       "аукцион",
		Sum:         "150 000",
	},
}
