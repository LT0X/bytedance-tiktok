package models

// ResponseStatus 响应状态码和信息
type ResponseStatus struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
