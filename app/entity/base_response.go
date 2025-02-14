package entity

type BaseResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}
