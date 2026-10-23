package http_server

import "time"

type Worker struct {
	FullName string `json:"FullName"`
	Position string `json:"Position"`
}

type ErrorRequest struct {
	Message   string    `json:"message"`
	TimeError time.Time `json:"time_error"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}
