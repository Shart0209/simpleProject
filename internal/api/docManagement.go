package api

import (
	"encoding/json"
	"fmt"
	"simpleProject/pkg/db"
	"simpleProject/pkg/model"
	"time"
)

func (s *service) GetAll() (map[int]model.DocumentManagement, error) {
	return db.DataBaseTest, nil
}

func (s *service) GetID(id int) (model.DocumentManagement, error) {
	data, ok := db.DataBaseTest[id]
	if ok {
		return data, nil
	}

	s.logger.Fatal().Msg("error get ID")
	return model.DocumentManagement{}, fmt.Errorf("from id=%d not found", id)
}

func (s *service) Add(arr []byte) (model.DocumentManagement, error) {
	var res model.DocumentManagement

	err := json.Unmarshal(arr, &res)
	if err != nil {
		s.logger.Error().Err(err).Msg("bad unmarshal JSON")
		return model.DocumentManagement{}, err
	}
	//add to map DB
	db.DataBaseTest[len(db.DataBaseTest)+1] = res

	return res, nil
}

func (s *service) UpdateID(id int) (string, error) {
	return "UpdateDocs", nil
}

func (s *service) DeleteID(id int) (string, error) {
	return "DeleteDocs", nil
}

func (s *service) DeleteALL(id int) (string, error) {
	return "DeleteDocs", nil
}

func getTime(i interface{}) time.Time {
	date := fmt.Sprintf("%v", i)
	t, _ := time.Parse(time.RFC3339, date)

	return t
}
