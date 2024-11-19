package models

type LogInfo struct {
	LogDefault
	ApplicationName   string `json:"application_name"`
	ClientRecognition string `json:"client_recognition"`
	HostName          string `json:"host_name"`
	HttpStatus        int    `json:"http_status"`
	Method            string `json:"method"`
	Msg               any    `json:"msg"`
	RequestBody       any    `json:"request_body"`
	ResponseBody      any    `json:"response_body"`
	ResposeTime       int    `json:"respose_time"`
	ServerIp          string `json:"server_ip"`
	StatusCode        string `json:"status_code"`
	Username          string `json:"username"`
}
