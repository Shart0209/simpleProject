package model

import (
	"mime/multipart"
	"time"
)

type File struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Path string `json:"path"`
}

type Files struct {
	File
	Files *multipart.FileHeader
}

type DocumentManagement struct {
	Id          int                    `form:"id" binding:"omitempty" db:"contract_id"`
	Title       string                 `form:"title" binding:"required" db:"title"`
	Number      string                 `form:"number" binding:"required" db:"contr_number"`
	Date        time.Time              `form:"date" binding:"required" time_format:"2006-01-02" db:"contr_date"`
	Category    string                 `form:"category" binding:"required" db:"category_name"`
	Price       float64                `form:"price" binding:"required" db:"price"`
	StartDate   time.Time              `form:"start_date" binding:"required" time_format:"2006-01-02" db:"start_date"`
	EndDate     time.Time              `form:"end_date" binding:"required" time_format:"2006-01-02" db:"end_date"`
	Description string                 `form:"description" binding:"omitempty" db:"description"`
	Distributor string                 `form:"distributor" binding:"omitempty" db:"company_name"`
	City        string                 `form:"city" binding:"omitempty" db:"company_city"`
	CreatedAt   time.Time              `form:"created_at" binding:"omitempty" db:"created_at"`
	UpdateAt    time.Time              `form:"update_at" binding:"omitempty" db:"update_at"`
	Files       map[string]interface{} `json:"files" db:"files"`
}

type BindForm struct {
	BindFiles     []*multipart.FileHeader `form:"files" binding:"omitempty"`
	DocManagement *DocumentManagement
}
