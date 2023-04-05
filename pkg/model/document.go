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

type DocsAttrs struct {
	ID          int                    `form:"id" json:"id" binding:"omitempty" db:"contract_id"`
	Title       string                 `form:"title" json:"title" binding:"required" db:"title"`
	Number      string                 `form:"number" json:"number" binding:"required" db:"contr_number"`
	Date        time.Time              `form:"date" json:"date" binding:"required" time_format:"2006-01-02" db:"contr_date"`
	Category    string                 `form:"category" json:"category" binding:"required" db:"category_name"`
	Price       float64                `form:"price" json:"price" binding:"required" db:"price"`
	StartDate   time.Time              `form:"start_date" json:"start_date" binding:"required" time_format:"2006-01-02" db:"start_date"`
	EndDate     time.Time              `form:"end_date" json:"end_date" binding:"required" time_format:"2006-01-02" db:"end_date"`
	Description string                 `form:"description" json:"description" binding:"omitempty" db:"description"`
	Distributor string                 `form:"distributor" json:"distributor" binding:"omitempty" db:"distributor_name"`
	Status      bool                   `form:"status" json:"status" binding:"omitempty" db:"status"`
	CreatedAt   time.Time              `form:"created_at" json:"created_at" binding:"omitempty" db:"created_at"`
	UpdateAt    time.Time              `form:"update_at" json:"update_at" binding:"omitempty" db:"update_at"`
	AttrFiles   map[string]interface{} `json:"files" db:"files"`
}

type BindForm struct {
	BindFiles     []*multipart.FileHeader `form:"files" binding:"omitempty"`
	DocManagement *DocsAttrs
}

type Sps struct {
	Category    []map[string]interface{}
	Distributor []map[string]interface{}
}

func (b *BindForm) Verify() bool {
	return true
}
