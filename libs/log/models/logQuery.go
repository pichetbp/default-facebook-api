package models

type LogQuery struct {
	LogDefault
	RawQuery     string `json:"raw_query"`
	ExecuteTime  string `json:"execute_time"`
	RowsEffected int    `json:"rows_effected"`
	Error        string `json:"error"`
}
