package model

import (
	"mime/multipart"
	"time"
)

type FileName struct {
	name string
	size int
}

type DocumentManagement struct {
	Id          uint      `form:"id" binding:"required" store:"id"`
	Name        string    `form:"name" binding:"required" store:"name"`
	Date        time.Time `form:"date" binding:"required" time_format:"2006-01-02" store:"date"`             //дата заключения ГК
	Description string    `form:"description" binding:"omitempty" store:"description"`                       //примечание
	FileName    *FileName `form:"filename" binding:"omitempty" store:"file_name"`                            //UUID
	StartDate   time.Time `form:"start_date" binding:"required" time_format:"2006-01-02" store:"start_date"` //срок действия ГК - начало
	EndDate     time.Time `form:"end_date" binding:"required" time_format:"2006-01-02" store:"end_date"`     //срок действия ГК - конец
	Distributor string    `form:"distributor" binding:"omitempty" store:"distributor"`                       //поставщик услуг
	Method      string    `form:"method" binding:"required" store:"method"`                                  //метод проведения закупки
	Price       float64   `form:"sum" binding:"required" store:"price"`                                      //сумма ГК
	CreatedAt   time.Time `form:"created_at" binding:"required" store:"created_at"`
}

type BindForm struct {
	Files         []*multipart.FileHeader `form:"files" binding:"omitempty"`
	DocManagement *DocumentManagement
}
