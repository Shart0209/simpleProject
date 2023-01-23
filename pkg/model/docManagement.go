package model

import "time"

type DocumentManagement struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`        //дата заключения ГК
	Description string    `json:"description"` //примечание
	File        []byte    `json:"file"`        //UUID
	DateIn      time.Time `json:"date_in"`     //срок действия ГК - начало
	DateOut     time.Time `json:"date_out"`    //срок действия ГК - конец
	Distributor string    `json:"distributor"` //поставщик услуг
	Method      string    `json:"method"`      //метод проведения закупки
	Sum         float64   `json:"sum"`         //сумма ГК
}
