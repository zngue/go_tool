package http

type Request struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
