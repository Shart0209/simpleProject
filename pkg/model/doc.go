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
	ID          int         `form:"id" json:"id" binding:"omitempty" db:"contract_id"`
	Title       string      `form:"title" json:"title" binding:"required" db:"title"`
	Number      string      `form:"number" json:"number" binding:"required" db:"numb"`
	Price       float64     `form:"price" json:"price" binding:"required" db:"price"`
	Date        time.Time   `form:"date" json:"date" binding:"required" time_format:"2006-01-02" db:"date"`
	StartDate   time.Time   `form:"start_date" json:"start_date" binding:"required" time_format:"2006-01-02" db:"start_date"`
	EndDate     time.Time   `form:"end_date" json:"end_date" binding:"required" time_format:"2006-01-02" db:"end_date"`
	Description string      `form:"description" json:"description" binding:"omitempty" db:"description"`
	Category    int         `form:"category" json:"category" binding:"required" db:"category"`
	Supplier    int         `form:"supplier" json:"supplier" binding:"required" db:"suppliers"`
	Status      bool        `form:"status" json:"status" binding:"required" db:"status"`
	Group       int         `form:"group" json:"group" binding:"required" db:"сgroups"`
	Author      string      `form:"author" json:"author" binding:"required" db:"author"`
	CreatedAt   time.Time   `form:"created_at" json:"created_at" binding:"omitempty" db:"created_at"`
	UpdateAt    time.Time   `form:"update_at" json:"update_at" binding:"omitempty" db:"update_at"`
	AttrFiles   interface{} `json:"files" db:"files"`
}

type UpdDoc struct {
	ID          int       `form:"id" json:"id" binding:"omitempty"`
	Title       string    `form:"title" json:"title" binding:"omitempty"`
	Number      string    `form:"number" json:"number" binding:"omitempty"`
	Price       float64   `form:"price" json:"price" binding:"omitempty"`
	Date        time.Time `form:"date" json:"date" binding:"omitempty" time_format:"2006-01-02"`
	StartDate   time.Time `form:"start_date" json:"start_date" binding:"omitempty" time_format:"2006-01-02"`
	EndDate     time.Time `form:"end_date" json:"end_date" binding:"omitempty" time_format:"2006-01-02"`
	Description string    `form:"description" json:"description" binding:"omitempty"`
	Status      bool      `form:"status" json:"status" binding:"omitempty"`
	CategoryID  int       `form:"category" json:"category" binding:"omitempty"`
	AuthorID    int       `form:"author" json:"author" binding:"omitempty"`
	SupplierID  int       `form:"supplier" json:"supplier" binding:"omitempty"`
	GroupsID    int       `form:"group" json:"group" binding:"omitempty"`
}

type BindForm struct {
	BindFiles []*multipart.FileHeader `form:"files" binding:"omitempty"`
	Docs      *UpdDoc
}

type Sps struct {
	Category  []map[string]interface{}
	Suppliers []map[string]interface{}
	Groups    []map[string]interface{}
}
