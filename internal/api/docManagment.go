package api

import (
	"errors"
	"fmt"
	"strings"
)

func (s *service) Doc(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("bad name")
	}
	return fmt.Sprintf("Name %s", name), nil
}

func (s *service) ListDocs() {}

func (s *service) AddDocs() {}

func (s *service) DeleteDocs() {}

func (s *service) GetDocs(id int) {}

func (s *service) UpdateDocs(id int) {}
