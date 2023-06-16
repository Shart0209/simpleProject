package model

import (
	"mime/multipart"
	"time"
)

type File struct {
	ID   string `json:"id"`
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
	Number      string                 `form:"number" json:"number" binding:"required" db:"numb"`
	Price       float64                `form:"price" json:"price" binding:"required" db:"price"`
	Date        time.Time              `form:"date" json:"date" binding:"required" time_format:"2006-01-02" db:"date"`
	StartDate   time.Time              `form:"start_date" json:"start_date" binding:"required" time_format:"2006-01-02" db:"start_date"`
	EndDate     time.Time              `form:"end_date" json:"end_date" binding:"required" time_format:"2006-01-02" db:"end_date"`
	Description string                 `form:"description" json:"description" binding:"omitempty" db:"description"`
	Category    int                    `form:"category" json:"category" binding:"required" db:"category"`
	Supplier    int                    `form:"supplier" json:"supplier" binding:"required" db:"suppliers"`
	Status      bool                   `form:"status" json:"status" binding:"required" db:"status"`
	Group       int                    `form:"group" json:"group" binding:"required" db:"сgroups"`
	CreatedAt   time.Time              `form:"created_at" json:"created_at" binding:"omitempty" db:"created_at"`
	UpdateAt    time.Time              `form:"update_at" json:"update_at" binding:"omitempty" db:"update_at"`
	AttrFiles   map[string]interface{} `json:"files" db:"files"`
}

type UpdDoc struct {
	Title       string    `json:"title"`
	Number      string    `json:"number"`
	Price       float64   `json:"price"`
	Date        time.Time `json:"date"`
	Start_date  time.Time `json:"start_date"`
	End_date    time.Time `json:"end_date"`
	Description string    `json:"description"`
	Category_id int       `json:"category"`
	Status      bool      `json:"status"`
	Supplier_id int       `json:"supplier"`
	C_groups_id int       `json:"group"`
}

type BindForm struct {
	BindFiles []*multipart.FileHeader `form:"files" binding:"omitempty"`
	Docs      *DocsAttrs
}

type Sps struct {
	Category  []map[string]interface{}
	Suppliers []map[string]interface{}
	Groups    []map[string]interface{}
}
