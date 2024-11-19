package models

type LogDebug struct {
	LogDefault
	Msg any `json:"msg"`
}
