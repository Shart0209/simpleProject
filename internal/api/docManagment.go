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

//TODO create CRUD method
//create
//read
//delete
//update
