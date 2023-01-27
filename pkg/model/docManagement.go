package model

import (
	"mime/multipart"
	"time"
)

type DocumentManagement struct {
	Id          uint                         `form:"id,required,omitempty"`
	Name        string                       `form:"name,required,omitempty"`
	Date        time.Time                    `form:"date,required"`          //дата заключения ГК
	Description string                       `form:"description,omitempty"`  //примечание
	FileName    map[string]map[string]string `form:"filename,omitempty"`     //UUID
	StartDate   time.Time                    `form:"start_date,omitempty"`   //срок действия ГК - начало
	EndDate     time.Time                    `form:"end_date,omitempty"`     //срок действия ГК - конец
	Distributor string                       `form:"distributor,omitempty"`  //поставщик услуг
	Method      string                       `form:"method,omitempty"`       //метод проведения закупки
	Sum         float64                      `form:"sum,required,omitempty"` //сумма ГК
}

type BindForm struct {
	Files []*multipart.FileHeader `form:"files,omitempty"`
	*DocumentManagement
}
