// package model

// type WebResponse struct {
// 	Code   int         `json:"code"`
// 	Status string      `json:"status"`
// 	Data   interface{} `json:"data"`
// }

// func ResponseToClient(code int, status string, data interface{}) WebResponse {
// 	return WebResponse{
// 		Code:   code,
// 		Status: status,
// 		Data:   data,
// 	}
// }

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
