package externalserver

import "fmt"

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"errors,omitempty"`
}

type LoginRequest struct {
	Login    string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func testCheckAuth(login string, password string) error {

	if login != "test" || password != "test" {
		return fmt.Errorf("authentication failed")
	}
	return nil
}
