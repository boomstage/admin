package model

const (
	CodeOK           = 0
	CodeServerError  = 1
	CodeParamInvalid = 2
	CodeUnAuthorized = 3
	CodeUserBaned    = 4 // 被封禁
)

type UserSource int

const (
	UserSourceApp   UserSource = 0 // 应用
	UserSourceAdmin UserSource = 1 // 管理后台
	UserSourceSvc   UserSource = 2 // 内部服务
)

type BaseResponse struct {
	Code       int         `json:"code"`
	Toast      string      `json:"toast,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	ServerTime int64       `json:"server_time"`
}
