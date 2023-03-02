package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"simpleProject/pkg/model"
	"strconv"
)

func (s *service) GetAll() (*[]model.DocumentManagement, error) {
	repo := s.store.repo

	data := &[]model.DocumentManagement{}

	query := `SELECT
	contract_id, title, contr_number,
	contr_date, price, start_date, end_date,
	description, created_at, company_name, company_city,
	category_name
	FROM contracts
	JOIN distributors USING (distributor_id)
	JOIN categories USING (category_id)
	JOIN authors USING (author_id)`

	// TODO flag:
	//	- true: get all
	// 	- false: get by ID
	err := repo.Get(data, query, true)
	if err != nil {
		s.logger.Error().Err(err).Msg("errors while executing query repo")
		return nil, err
	}

	query = `SELECT file_name, file_size FROM files WHERE contract_id=$1`

	tmp := make([]model.DocumentManagement, 0, len(*data))

	for _, item := range *data {
		files := &[]model.File{}

		err = repo.Get(files, query, true, item.Id)
		if err != nil {
			s.logger.Error().Err(err).Msg("errors while executing query repo files")
			return nil, err
		}
		item.Files = append(item.Files, *files...)
		tmp = append(tmp, item)
	}

	return &tmp, nil
}

func (s *service) GetByID(id int) (*model.DocumentManagement, error) {

	repo := s.store.repo

	data := &model.DocumentManagement{}

	query := `SELECT
	contract_id, title, contr_number,
	contr_date, price, start_date, end_date,
	description, created_at, company_name, company_city,
	category_name
	FROM contracts
	JOIN distributors USING (distributor_id)
	JOIN categories USING (category_id)
	JOIN authors USING (author_id)
	WHERE contract_id=$1`

	err := repo.Get(data, query, false, id)
	if err != nil {
		s.logger.Error().Err(err).Msg("errors while executing query get ID")
		return nil, err
	}

	files := &[]model.File{}
	query = `SELECT file_name, file_size FROM files WHERE contract_id=$1`
	err = repo.Get(files, query, true, data.Id)
	if err != nil {
		s.logger.Error().Err(err).Msg("errors while executing query get files ID")
		return nil, err
	}

	data.Files = append(data.Files, *files...)

	return data, nil
}

func (s *service) Create(bindForm *model.BindForm) error {

	if files := bindForm.BindFiles; len(files) != 0 {
		arrNameFile, err := s.GetNameAndSaveFile(files)
		if err != nil {
			return err
		}

		fmt.Println(arrNameFile)
	}
	return nil

}

func (s *service) Update(id int, bindForm *model.BindForm) error {

	//if _, ok := db.DataBaseTest[id]; !ok {
	//	return fmt.Errorf("not found")
	//}
	//
	//if files := bindForm.Files; len(files) != 0 {
	//	arrNameFile, err := s.GetNameAndSaveFile(files)
	//	if err != nil {
	//		return err
	//	}
	//	//TODO падает при поллучении одного поля files
	//	bindForm.FileName = *arrNameFile
	//}
	//db.DataBaseTest[id] = *bindForm.DocumentManagement

	return nil
}

func (s *service) Delete(id int) (pgconn.CommandTag, error) {

	repo := s.store.repo

	query := `DELETE FROM contracts	WHERE contract_id = $1`

	tag, err := repo.Exec(query, id)

	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("failed to delete by id=%d", id)
	}

	return tag, nil
}

func (s *service) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err

}

func (s *service) GetNameAndSaveFile(files []*multipart.FileHeader) ([]map[string]string, error) {

	arrNameFile := make([]map[string]string, 0)

	baseDir, err := filepath.Abs(s.cfg.FilesFolder)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("path folder to ./upload not found")
	}

	fileID := uuid.New().String()

	for i, file := range files {
		fName := filepath.Base(file.Filename)
		filename := fileID + "-" + strconv.Itoa(i) + filepath.Ext(fName)

		tmp := map[string]string{
			"name": filename,
			"size": strconv.Itoa(int(file.Size)),
		}
		arrNameFile = append(arrNameFile, tmp)

		filePath := filepath.Join(baseDir, filename)
		if err := s.SaveUploadedFile(file, filePath); err != nil {
			s.logger.Error().Err(err).Send()
			return nil, fmt.Errorf("upload file err: %s", err.Error())
		}
	}
	return arrNameFile, nil
}
