package exception

type ErrorRes struct {
	Code    string `json:"code"`
	TraceId string `json:"trace_id"`
	Message string `json:"message"`
}
