package model

import (
	"mime/multipart"
	"time"
)

type DocumentManagement struct {
	Id          uint                         `form:"id" binding:"required" db:"id"`
	Name        string                       `form:"name" binding:"required" db:"name"`
	Date        time.Time                    `form:"date" binding:"required" time_format:"2006-01-02" db:"date"`             //дата заключения ГК
	Description string                       `form:"description" binding:"omitempty" db:"description"`                       //примечание
	FileName    map[string]map[string]string `form:"filename" binding:"omitempty" db:"file_name"`                            //UUID
	StartDate   time.Time                    `form:"start_date" binding:"required" time_format:"2006-01-02" db:"start_date"` //срок действия ГК - начало
	EndDate     time.Time                    `form:"end_date" binding:"required" time_format:"2006-01-02" db:"end_date"`     //срок действия ГК - конец
	Distributor string                       `form:"distributor" binding:"omitempty" db:"distributor"`                       //поставщик услуг
	Method      string                       `form:"method" binding:"required" db:"method"`                                  //метод проведения закупки
	Sum         float64                      `form:"sum" binding:"required" db:"sum"`                                        //сумма ГК
}

type BindForm struct {
	Files []*multipart.FileHeader `form:"files" binding:"omitempty"`
	*DocumentManagement
}
