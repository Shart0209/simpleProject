package api

import (
	"encoding/json"
	"fmt"
	"simpleProject/pkg/constants"
	"simpleProject/pkg/model"
	"simpleProject/pkg/util"
	"time"
)

func (s *service) GetAll() ([]*model.DocumentManagement, error) {
	start := time.Now()

	var data []*model.DocumentManagement

	query := fmt.Sprintf(`SELECT %s 
	FROM contracts
	JOIN distributors USING (distributor_id)
	JOIN categories USING (category_id)
	LEFT JOIN authors USING (author_id)
	ORDER BY contract_id DESC;`, constants.Repo.GetColumns)

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := s.store.repo.Get(&data, query, true)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	}

	duration := time.Since(start)
	fmt.Println(duration) //15.872689ms -> 626.186µs

	return data, nil
}

func (s *service) GetByID(id uint64) (*model.DocumentManagement, error) {

	start := time.Now()

	var data model.DocumentManagement

	query := fmt.Sprintf(`SELECT %s 
	FROM contracts
	JOIN distributors USING (distributor_id)
	JOIN categories USING (category_id)
	LEFT JOIN authors USING (author_id)
	WHERE contract_id=$1
	ORDER BY contract_id DESC;`, constants.Repo.GetColumns)

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := s.store.repo.Get(&data, query, false, int(id))
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("ID not found id=%v", id)
	}

	duration := time.Since(start)
	fmt.Println(duration) //

	return &data, nil

}

func (s *service) Create(bindForm *model.BindForm) error {
	start := time.Now()

	var jsonFiles []byte

	data := []model.Files{}

	if len(bindForm.BindFiles) != 0 {

		err := util.ParserBindForm(bindForm.BindFiles, s.cfg.FilesFolder, &data)
		if err != nil {
			return fmt.Errorf("failed to parse bind form")
		}

		var res []model.File
		for _, item := range data {
			tmp := item.File
			res = append(res, tmp)
		}
		attr := make(map[string]interface{}, 1)
		attr["attr"] = res

		jsonFiles, err = json.Marshal(attr)
		if err != nil {
			return err
		}

	}

	query := fmt.Sprintf("INSERT INTO contracts	(%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		constants.Repo.CreateColums)

	err := s.store.repo.InsertOne(nil, query,
		bindForm.DocManagement.Title,
		bindForm.DocManagement.Number,
		bindForm.DocManagement.Date,
		bindForm.DocManagement.Category,
		bindForm.DocManagement.Price,
		bindForm.DocManagement.StartDate,
		bindForm.DocManagement.EndDate,
		bindForm.DocManagement.Description,
		bindForm.DocManagement.Distributor,
		jsonFiles,
	)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return fmt.Errorf("failed to create")
	}

	for _, item := range data {
		if err := util.SaveUploadedFile(item.Files, &item.File.Path); err != nil {
			s.logger.Error().Err(err).Send()
			return fmt.Errorf("upload file err: %s", err.Error())
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)

	return nil
}

func (s *service) Update(id int, bindForm *model.BindForm) error {

	return nil
}

func (s *service) Delete(id uint64) error {
	start := time.Now()

	//delete records by id
	query := `DELETE FROM contracts WHERE contract_id=$1 RETURNING files;`

	var data model.DocumentManagement
	err := s.store.repo.InsertOne(&data.AttrFiles, query, int(id))
	if err != nil {
		s.logger.Error().Err(err).Send()
		return fmt.Errorf("failed %w", err)
	}

	var items = data.AttrFiles["attr"].([]interface{})
	for _, item := range items {
		res := item.(map[string]interface{})
		path := res["path"].(string)
		if err := util.DeleteFile(&path); err != nil {
			return fmt.Errorf("failed to delete file %w", err)
		}
	}

	duration := time.Since(start)
	fmt.Println(duration) //15.872689ms -> 626.186µs

	return nil
}
