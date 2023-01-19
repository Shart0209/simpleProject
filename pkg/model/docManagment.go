package model

type DocumentManagment struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`        //дата заключения ГК
	Description string `json:"description"` //примечание
	File        string `json:"file"`        //UUID
	DateIn      string `json:"date_in"`     //срок действия ГК - начало
	DateOut     string `json:"date_out"`    //срок действия ГК - конец
	Distributor string `json:"distributor"` //поставщик услуг
	Metod       string `json:"metod"`       //метод проведения закупки
	Sum         string `json:"sum"`         //сумма ГК
}
