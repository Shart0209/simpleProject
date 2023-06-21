package api

import (
	"encoding/json"
	"fmt"
	"os"
	filepath "path/filepath"
	"simpleProject/pkg/constants"
	"simpleProject/pkg/model"
	"simpleProject/pkg/util"
	"strings"
	"time"
)

func (s *service) GetAll() ([]*model.DocsAttrs, error) {
	start := time.Now()

	var data []*model.DocsAttrs

	query := fmt.Sprintf(`SELECT %s 
	FROM commons
	JOIN contracts USING (contract_id)
	JOIN suppliers USING (supplier_id)
    JOIN authors USING (author_id)
	JOIN categories USING (category_id)
	JOIN c_groups USING (c_groups_id)
	ORDER BY contract_id DESC;`, constants.Repo.Colum)

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := s.store.repo.Get(&data, query, true)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("data is empty")
	}

	duration := time.Since(start)
	fmt.Println(duration)

	return data, nil
}

func (s *service) GetByID(ID uint64) (*model.DocsAttrs, error) {

	start := time.Now()

	var data model.DocsAttrs

	query := fmt.Sprintf(`SELECT %s 
	FROM commons
	JOIN authors USING (author_id)
	JOIN contracts USING (contract_id)
	JOIN suppliers USING (supplier_id)
	JOIN categories USING (category_id)
	JOIN c_groups USING (c_groups_id)
	WHERE commons.contract_id=$1;`, constants.Repo.Colum)

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := s.store.repo.Get(&data, query, false, int(ID))
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("ID not found ID=%v", ID)
	}

	duration := time.Since(start)
	fmt.Println(duration) //

	return &data, nil

}

func (s *service) GetSps() (*model.Sps, error) {

	var out [][]map[string]interface{}

	query := []string{
		"SELECT category_id AS id, categories.name AS name FROM categories",
		"SELECT supplier_id AS id, suppliers.name AS name FROM suppliers",
		"SELECT c_groups_id AS id, c_groups.name AS name FROM c_groups"}
	for i := range query {
		var data []map[string]interface{}
		err := s.store.repo.Get(&data, query[i], true)
		if err != nil {
			s.logger.Error().Err(err).Send()
		}
		out = append(out, data)
	}

	sps := model.Sps{
		Category:  out[0],
		Suppliers: out[1],
		Groups:    out[2],
	}

	return &sps, nil
}

func (s *service) Create(bindForm *model.BindForm) error {
	start := time.Now()

	var jsonFiles []byte

	data := []model.Files{}

	res := util.ReplaceNameFolder(&bindForm.Docs.Number)

	folderPath := filepath.Join(s.cfg.FilesFolder, res)
	if err := util.CreateFolder(&folderPath); err != nil {
		s.baseLogger.Error().Err(err).Send()
		return err
	}

	if len(bindForm.BindFiles) != 0 {

		err := util.ParserBindForm(bindForm.BindFiles, folderPath, &data)
		if err != nil {
			return fmt.Errorf("failed to parse bind form")
		}

		var res []model.File
		for _, item := range data {
			res = append(res, item.File)
		}

		jsonFiles, err = json.Marshal(res)
		if err != nil {
			return err
		}
	}

	query := `WITH new_contract as  (
	INSERT INTO contracts (title, numb, price, date, start_date, end_date, description, files)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING contract_id
	)
	INSERT INTO commons
	SELECT contract_id, $9, $10, $11, $12 from new_contract`
	err := s.store.repo.InsertOne(nil, query,
		bindForm.Docs.Title,
		bindForm.Docs.Number,
		bindForm.Docs.Price,
		bindForm.Docs.Date,
		bindForm.Docs.StartDate,
		bindForm.Docs.EndDate,
		bindForm.Docs.Description,
		jsonFiles,
		bindForm.Docs.SupplierID,
		bindForm.Docs.CategoryID,
		bindForm.Docs.GroupsID,
		bindForm.Docs.AuthorID,
	)
	if err != nil {
		s.logger.Error().Err(err).Send()
		if err := util.DeleteFile(&folderPath); err != nil {
			return fmt.Errorf("failed to create and delete folder")
		}
		return fmt.Errorf("failed to create")
	}

	for _, item := range data {
		if err := util.SaveUploadedFile(item.Files, &item.File.ID, &item.File.Path); err != nil {
			s.logger.Error().Err(err).Send()
			return fmt.Errorf("upload file err: %s", err.Error())
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)

	return nil
}

func (s *service) Update(bindForm *model.BindForm) error {

	oldItem, err := s.GetByID(uint64(bindForm.Docs.ID))
	if err != nil {
		return err
	}

	if oldItem.Number != bindForm.Docs.Number {
		newRes := util.ReplaceNameFolder(&bindForm.Docs.Number)
		newFolderPath := filepath.Join(s.cfg.FilesFolder, newRes)

		oldRes := util.ReplaceNameFolder(&oldItem.Number)
		oldFolderPath := filepath.Join(s.cfg.FilesFolder, oldRes)
		err := util.RenameFolder(&oldFolderPath, &newFolderPath)
		if err != nil {
			return err
		}

		tmp := len(oldItem.AttrFiles)
		_ = tmp
		//var tmp []map[string]interface{}
		//coutnFilesToJson := json.Unmarshal(oldItem.AttrFiles, &tmp)
		//query := `UPDATE contracts
		//		SET files=jsonb_set(files, '{$1, path}', '"$2"', false)
		//		WHERE contract_id=$3`

		//err := s.store.repo.InsertOne(nil, query)
	}

	//query := `WITH upd_contract as (
	//UPDATE contracts
	//SET title='blablabla', numb='12345-576'
	//WHERE contract_id =3
	//RETURNING contract_id as id
	//)
	//UPDATE commons
	//SET status=false
	//WHERE contract_id=(SELECT id from upd_contract)`

	return nil
}

func (s *service) Delete(ID uint64) error {
	start := time.Now()

	//delete files by ID from folder
	query := `DELETE FROM contracts WHERE contract_id=$1 RETURNING numb;`

	var data string
	err := s.store.repo.InsertOne(&data, query, int(ID))
	if err != nil {
		s.logger.Error().Err(err).Send()
		return fmt.Errorf("failed %w", err)
	}

	baseDir, err := filepath.Abs(s.cfg.FilesFolder)
	if err != nil {
		return fmt.Errorf("path folder to ./upload/%s not found", data)
	}
	res := util.ReplaceNameFolder(&data)
	pathDir := filepath.Join(baseDir, res)

	if err := util.DeleteFile(&pathDir); err != nil {
		return fmt.Errorf("failed to delete folder %w", err)
	}
	//var pathDir string
	//if data.AttrFiles != nil {
	//	var items = data.AttrFiles["attr"].([]interface{})
	//	pathDir = filepath.Dir(items[0].(map[string]interface{})["path"].(string))
	//	for _, item := range items {
	//		res := item.(map[string]interface{})
	//		path := res["path"].(string)
	//		if err := util.DeleteFile(&path); err != nil {
	//			return fmt.Errorf("failed to delete file %w", err)
	//		}
	//	}

	duration := time.Since(start)
	fmt.Println(duration)

	return nil
}

func (s *service) GetFilePath(fileName string, ID string) (string, error) {

	res := strings.Split(fileName, ".")[0]
	idx := res[len(res)-1:]

	query := fmt.Sprintf("SELECT files::json->'attr'->%s->'path' FROM contracts WHERE contract_id = %s", idx, ID)

	var data interface{}
	err := s.store.repo.Get(&data, query, false)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return "", err
	}

	path := filepath.Join(data.(string), fileName)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			s.logger.Error().Err(err).Msg("file does not exist")
			return "", err
		}
	}

	return data.(string), nil
}
