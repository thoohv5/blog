package common

// 响应数据体
type ResponseEntity struct {
	// 错误码
	ErrorCode int `json:"errcode"`
	// 响应消息
	Message string `json:"errmsg"`
	// 响应数据
	Data interface{} `json:"data"`
}
