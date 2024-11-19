package models

type LogError struct {
	LogDefault
	Error string `json:"error"`
}
