package models

import (
	"time"
)

// class
type LogDefault struct {
	Account       string    `json:"account"`
	FunctionName  string    `json:"function_name"`
	Level         string    `json:"level"`
	TimeStamp     time.Time `json:"time_stamp"`
	TransactionId string    `json:"transaction_id"`
	FileLocation  string    `json:"file_location"`
	Type          string    `json:"type"`
}
