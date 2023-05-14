package response

type SuccessResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}
