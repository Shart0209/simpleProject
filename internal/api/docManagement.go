package api

import (
	"encoding/json"
	"fmt"
	"simpleProject/pkg/db"
	"simpleProject/pkg/model"
)

func (s *service) GetAll() (map[int]model.DocumentManagement, error) {
	return db.DataBaseTest, nil
}

func (s *service) GetID(id int) (model.DocumentManagement, error) {
	if data, ok := db.DataBaseTest[id]; ok {
		return data, nil
	}

	s.logger.Error().Msg("error while executing query get ID")
	return model.DocumentManagement{}, fmt.Errorf("could not found record by id=%d", id)
}

func (s *service) Add(arr *[]byte) (model.DocumentManagement, error) {
	var res model.DocumentManagement

	err := json.Unmarshal(*arr, &res)
	if err != nil {
		s.logger.Error().Err(err).Msg("error unmarshal JSON from ctx body for add")
		return model.DocumentManagement{}, err
	}
	//add to map DB
	db.DataBaseTest[len(db.DataBaseTest)+1] = res

	return res, nil
}

func (s *service) UpdateID(id int, arr *[]byte) (model.DocumentManagement, error) {
	var res model.DocumentManagement

	if _, ok := db.DataBaseTest[id]; ok {
		err := json.Unmarshal(*arr, &res)
		if err != nil {
			s.logger.Error().Err(err).Msg("error unmarshal JSON from ctx body for update")
			return model.DocumentManagement{}, err
		}

		//update to map DB
		db.DataBaseTest[id] = res
		return db.DataBaseTest[id], nil
	}

	return model.DocumentManagement{}, fmt.Errorf("failed to update record by id=%d", id)
}

func (s *service) DeleteID(id int) error {

	if _, ok := db.DataBaseTest[id]; ok {
		delete(db.DataBaseTest, id)
		return nil
	}

	return fmt.Errorf("failed to delete record by id=%d", id)
}

func (s *service) DeleteALL() {
	db.DataBaseTest = map[int]model.DocumentManagement{}
}
