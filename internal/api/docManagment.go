package api

import (
	"errors"
	"fmt"
	"strings"
)

func (s *service) Hello(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("bad name")
	}
	return fmt.Sprintf("Hello world %s", name), nil
}

//TODO create CRUD method
//create
//read
//delete
//update
