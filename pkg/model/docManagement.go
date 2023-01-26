package model

import (
	"time"
)

type DocumentManagement struct {
	Id          uint      `json:"id" binding:"required,omitempty"`
	Name        string    `json:"name" binding:"required,omitempty"`
	Date        time.Time `json:"date" binding:"required,omitempty"`        //дата заключения ГК
	Description string    `json:"description" binding:"required,omitempty"` //примечание
	FileName    []string  `json:"filename" binding:"required,omitempty"`    //UUID
	StartDate   time.Time `json:"start_date" binding:"required,omitempty"`  //срок действия ГК - начало
	EndDate     time.Time `json:"end_date" binding:"required,omitempty"`    //срок действия ГК - конец
	Distributor string    `json:"distributor" binding:"required,omitempty"` //поставщик услуг
	Method      string    `json:"method" binding:"required,omitempty"`      //метод проведения закупки
	Sum         float64   `json:"sum" binding:"required,omitempty"`         //сумма ГК
}
