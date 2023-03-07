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
	"strings"
	"time"
)

type files struct {
	mf   model.File
	file *multipart.FileHeader
}

func (s *service) GetAll() ([]*model.DocumentManagement, error) {
	start := time.Now()

	repo := s.store.repo
	var data []*model.DocumentManagement

	query := `SELECT c.contract_id,
    title,
    contr_number,
    contr_date,
    price,
    start_date,
    end_date,
    description,
    created_at,
    company_name,
    company_city,
    category_name,
    array_agg(file_id || ' ' || file_name || ' ' || file_size) arr_file
	FROM contracts c
      LEFT JOIN (SELECT contract_id, file_id, file_name, file_size
                 FROM files
                 GROUP BY file_id, file_name, file_size) t USING (contract_id)
      JOIN distributors USING (distributor_id)
      JOIN categories USING (category_id)
      LEFT JOIN authors USING (author_id)
	GROUP BY c.contract_id, title, contr_number, contr_date, price, start_date, end_date, description, created_at, company_name,
      company_city, category_name
	ORDER BY c.contract_id DESC;`

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := repo.Get(&data, query, true)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	}

	for _, item := range data {
		if item.Files[0] != nil {
			var tmpFiles []interface{}
			tmpFiles = append(tmpFiles, item.Files...)
			item.Files = nil

			for _, s := range tmpFiles {
				s := fmt.Sprintf("%v", s)
				st := strings.Split(s, " ")

				tmp := model.File{}
				tmp.Id, _ = strconv.Atoi(st[0])
				tmp.Name = st[1]
				tmp.Size, _ = strconv.Atoi(st[2])

				item.Files = append(item.Files, tmp)
			}
		}
	}

	duration := time.Since(start)
	fmt.Println(duration) //15.872689ms -> 626.186µs

	return data, nil
}

func (s *service) GetByID(id uint64) (*model.DocumentManagement, error) {

	start := time.Now()

	repo := s.store.repo
	var data model.DocumentManagement

	query := `SELECT c.contract_id,
    title,
    contr_number,
    contr_date,
    price,
    start_date,
    end_date,
    description,
    created_at,
    company_name,
    company_city,
    category_name,
    array_agg(file_id || ' ' || file_name || ' ' || file_size) arr_file
	FROM contracts c
      LEFT JOIN (SELECT contract_id, file_id, file_name, file_size
                 FROM files
                 GROUP BY file_id, file_name, file_size) t USING (contract_id)
      JOIN distributors USING (distributor_id)
      JOIN categories USING (category_id)
      LEFT JOIN authors USING (author_id)
	WHERE c.contract_id=$1
	GROUP BY c.contract_id, title, contr_number, 
	         contr_date, price, start_date, end_date, 
	         description, created_at, company_name, 
	         company_city, category_name
	ORDER BY c.contract_id DESC;`

	// TODO flag:
	//	- true: get ScanAll
	// 	- false: get ScanOne
	err := repo.Get(&data, query, false, id)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("ID not found id=%v", id)
	}

	if data.Files[0] != nil {
		var tmpFiles []interface{}
		tmpFiles = append(tmpFiles, data.Files...)
		data.Files = nil

		for _, s := range tmpFiles {
			s := fmt.Sprintf("%v", s)
			st := strings.Split(s, " ")

			tmp := model.File{}
			tmp.Id, _ = strconv.Atoi(st[0])
			tmp.Name = st[1]
			tmp.Size, _ = strconv.Atoi(st[2])

			data.Files = append(data.Files, tmp)
		}

	}

	duration := time.Since(start)
	fmt.Println(duration) //

	return &data, nil
}

func (s *service) Create(bindForm *model.BindForm) error {
	start := time.Now()

	repo := s.store.repo

	query := `
INSERT INTO contracts	(title, contr_number, contr_date, category_id, price, start_date, end_date,
description, distributor_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING contract_id`

	data := model.DocumentManagement{}
	err := repo.InsertOne(&data.Id, query,
		bindForm.DocManagement.Title,
		bindForm.DocManagement.Number,
		bindForm.DocManagement.Date,
		bindForm.DocManagement.Category,
		bindForm.DocManagement.Price,
		bindForm.DocManagement.StartDate,
		bindForm.DocManagement.EndDate,
		bindForm.DocManagement.Description,
		bindForm.DocManagement.Distributor,
	)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return fmt.Errorf("failed to create")
	}

	bindFiles := bindForm.BindFiles
	if len(bindFiles) != 0 {
		arrFiles, err := s.GetNameFiles(bindFiles)
		if err != nil {
			return err
		}

		query = `INSERT INTO files (file_name, file_size, file_path, contract_id) VALUES ($1, $2, $3, $4)`

		for _, item := range arrFiles {
			err := repo.InsertOne(nil, query,
				item.mf.Name,
				item.mf.Size,
				item.mf.Path,
				data.Id)
			if err != nil {
				return err
			}

		}
		if err := s.SaveUploadedFile(arrFiles); err != nil {
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

func (s *service) Delete(id uint64) (pgconn.CommandTag, error) {
	start := time.Now()

	repo := s.store.repo

	// get file_path by ID
	var data []*model.DocumentManagement
	query := `SELECT contract_id, array_agg(file_path) arr_file FROM files f WHERE contract_id=$1 GROUP BY contract_id`
	err := repo.Get(&data, query, true, int(id))
	if err != nil {
		s.logger.Error().Err(err).Send()
		return pgconn.CommandTag{}, err
	}

	if data == nil {
		return pgconn.CommandTag{}, fmt.Errorf("ID not found id=%v", id)
	}

	//delete records by id
	query = `DELETE FROM contracts WHERE contract_id=$1;`
	tag, err := repo.Exec(query, int(id))
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("failed to delete by id=%d", id)
	}

	//delete files to ./upload
	dt := data[0]
	if dt.Files[0] != nil {
		for _, items := range dt.Files {
			line := fmt.Sprintf("%v", items)
			lines := strings.Split(line, ",")
			for _, item := range lines {
				//checking that the file does not exist
				if _, err := os.Stat(item); err != nil {
					if os.IsNotExist(err) {
						fmt.Println("file does not exist")
						continue
					}
				}
				err := os.Remove(item)
				if err != nil {
					s.logger.Error().Err(err).Send()
					return pgconn.CommandTag{}, err
				}
			}
		}
	}

	duration := time.Since(start)
	fmt.Println(duration) //15.872689ms -> 626.186µs

	return tag, nil
}

func (s *service) GetNameFiles(bindFiles []*multipart.FileHeader) ([]*files, error) {

	arrFiles := make([]*files, 0, len(bindFiles))

	baseDir, err := filepath.Abs(s.cfg.FilesFolder)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return []*files{}, fmt.Errorf("path folder to ./upload not found")
	}

	fileID := uuid.New().String()

	for i, file := range bindFiles {
		fName := filepath.Base(file.Filename)
		filename := fileID + "-" + strconv.Itoa(i) + filepath.Ext(fName)

		filePath := filepath.Join(baseDir, filename)

		tmp := files{
			mf: model.File{
				Name: filename,
				Size: int(file.Size),
				Path: filePath,
			},
			file: file,
		}

		arrFiles = append(arrFiles, &tmp)
	}
	return arrFiles, nil
}

func (s *service) SaveUploadedFile(files []*files) error {

	for _, item := range files {
		src, err := item.file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		out, err := os.Create(item.mf.Path)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		if err != nil {
			return err
		}
	}
	return nil
}
