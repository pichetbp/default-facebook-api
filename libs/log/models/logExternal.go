package models

type LogExternal struct {
	LogDefault
	RequestMethod string `json:"request_method"`
	RequestUrl    string `json:"request_url"`
	RequestHeader any    `json:"request_header"`
	RequestBody   any    `json:"request_body"`
	HttpStatus    int    `json:"http_status"`
	ResponseBody  any    `json:"response_body"`
	ResponseTime  int    `json:"response_time"`
}
