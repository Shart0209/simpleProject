package api

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"simpleProject/pkg/db"
	"simpleProject/pkg/model"
	"strconv"
)

func (s *service) GetAll() (*map[int]model.DocumentManagement, error) {
	return &db.DataBaseTest, nil
}

func (s *service) GetByID(id int) (*model.DocumentManagement, error) {
	if data, ok := db.DataBaseTest[id]; ok {
		return &data, nil
	}

	s.logger.Error().Msg("errors while executing query get ID")
	return &model.DocumentManagement{}, fmt.Errorf("could not found record by id=%d", id)
}

func (s *service) Create(bindForm *model.BindForm) error {

	if files := bindForm.Files; len(files) != 0 {
		arrNameFile, err := s.GetNameAndSaveFile(files)
		if err != nil {
			return err
		}
		bindForm.DocumentManagement.FileName = *arrNameFile
	}

	db.DataBaseTest[len(db.DataBaseTest)+1] = *bindForm.DocumentManagement

	return nil
}

func (s *service) Update(id int, bindForm *model.BindForm) error {

	if _, ok := db.DataBaseTest[id]; !ok {
		return fmt.Errorf("not found")
	}

	if files := bindForm.Files; len(files) != 0 {
		arrNameFile, err := s.GetNameAndSaveFile(files)
		if err != nil {
			return err
		}
		//TODO падает при поллучении одного поля files
		bindForm.FileName = *arrNameFile
	}
	db.DataBaseTest[id] = *bindForm.DocumentManagement

	return nil
}

func (s *service) Delete(id int) error {

	if _, ok := db.DataBaseTest[id]; ok {
		delete(db.DataBaseTest, id)
		return nil
	}

	return fmt.Errorf("failed to delete record by id=%d", id)
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

func (s *service) GetNameAndSaveFile(files []*multipart.FileHeader) (*map[string]map[string]string, error) {

	arrNameFile := make(map[string]map[string]string, 0)

	baseDir, err := filepath.Abs(s.cfg.FilesFolder)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("path folder to ./upload not found")
	}

	fileID := uuid.New().String()

	for i, file := range files {
		fName := filepath.Base(file.Filename)
		filename := fileID + "-" + strconv.Itoa(i) + filepath.Ext(fName)
		arrNameFile[fName] = map[string]string{
			"name": filename,
			"size": strconv.Itoa(int(file.Size)),
		}

		filePath := filepath.Join(baseDir, filename)
		if err := s.SaveUploadedFile(file, filePath); err != nil {
			s.logger.Error().Err(err).Send()
			return nil, fmt.Errorf("upload file err: %s", err.Error())
		}
	}
	return &arrNameFile, nil
}
