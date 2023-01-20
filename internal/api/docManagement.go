package api

import (
	"fmt"
	"simpleProject/pkg/db"
	"simpleProject/pkg/model"
)

func (s *service) GetAll() (map[int]model.DocumentManagement, error) {
	return db.DataBaseTest, nil
}

func (s *service) GetID(id int) (string, error) {
	return fmt.Sprintf("GetDocs %v", id), nil
}

func (s *service) Add() (string, error) {
	return "AddDocs", nil
}

func (s *service) UpdateID(id int) (string, error) {
	return "UpdateDocs", nil
}

func (s *service) DeleteID(id int) (string, error) {
	return "DeleteDocs", nil
}
