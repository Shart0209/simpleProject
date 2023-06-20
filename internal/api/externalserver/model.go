package externalserver

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"errors,omitempty"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
