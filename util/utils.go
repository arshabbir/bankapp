package util

type ApiError struct {
	Statuscode int    `json:"statuscode"`
	Message    string `json:"message"`
}
