package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"simpleProject/pkg/db"
	"simpleProject/pkg/model"
)

func (s *service) getJSON(arr *[]byte) (*model.DocumentManagement, error) {
	var res model.DocumentManagement
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(*arr, &res); err != nil {
		s.logger.Error().Err(err).Msg("error decoding json response")
		return &model.DocumentManagement{}, err
	}
	return &res, nil
}

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

func (s *service) Add(arr *[]byte, files *[]string) (model.DocumentManagement, error) {

	res, err := s.getJSON(arr)
	if err != nil {
		return model.DocumentManagement{}, err
	}

	res.FileName = *files

	//add to map DB
	db.DataBaseTest[len(db.DataBaseTest)+1] = *res

	return *res, nil
}

func (s *service) UpdateID(id int, arr *[]byte) (model.DocumentManagement, error) {

	if _, ok := db.DataBaseTest[id]; ok {
		res, err := s.getJSON(arr)
		if err != nil {
			return model.DocumentManagement{}, err
		}

		//update to map DB
		db.DataBaseTest[id] = *res
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
