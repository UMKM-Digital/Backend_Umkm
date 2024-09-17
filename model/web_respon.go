package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Message string     `json:"message"`
	Data   interface{} `json:"data"`
}

func ResponseToClient(code int, status bool, message string, data interface{}) WebResponse {
	return WebResponse{
		Code:   code,
		Status: status,
		Message: message,
		Data:   data,
	}
}


type WebResponsepagi struct {
	Code       int           `json:"code"`
	Status     string        `json:"status"` // Change to string to match the response format
	Message    string        `json:"message"`
	Pagination Pagination    `json:"meta"`
	Data       interface{}   `json:"data"`
}

type Pagination struct {
	CurrentPage   int `json:"current_page"`
	NextPage      *int `json:"next_page"`
	PrevPage      *int `json:"prev_page"`
	TotalPages    int `json:"total_pages"`
	TotalRecords  int `json:"total_records"`
}

func ResponseToClientpagi(code int, status string, message string, pagination Pagination, data interface{}) WebResponsepagi {
	return WebResponsepagi{
		Code:       code,
		Status:     status,
		Message:    message,
		Pagination: pagination,
		Data:       data,
	}
}