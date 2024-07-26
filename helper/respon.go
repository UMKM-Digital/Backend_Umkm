package helper

type ErrorResponse struct {
    Code    int    `json:"code"`
    Status  string `json:"status"`
    Data    interface{} `json:"data"`
}