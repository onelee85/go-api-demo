package common

// Common
const (
	ErrCodeFail = iota
	ErrCodeOk
	ErrCodeHttp
	ErrCodeSystem
	ErrCodeNotFound
	ErrCodeDuplicated
	ErrCodeNoPermission
	ErrCodeInvalidParams
	ErrCodeInvalidVerifyCode
)

// Account
const (
	ErrCodeWrongPassword = 1000 + iota
)

var ErrMessages = map[int]string{
	ErrCodeFail:              "失败",
	ErrCodeOk:                "成功",
	ErrCodeHttp:              "请求错误",
	ErrCodeSystem:            "系统错误",
	ErrCodeNotFound:          "资源未找到",
	ErrCodeDuplicated:        "资源重复",
	ErrCodeNoPermission:      "没有权限",
	ErrCodeInvalidParams:     "参数错误",
	ErrCodeInvalidVerifyCode: "验证码错误",

	ErrCodeWrongPassword: "密码错误",
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult(code int, message string, data interface{}) (result *Result) {
	if message == "" {
		message = ErrMessages[code]
	}
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ResponseResult(code int, data interface{}) (result *Result) {
	return NewResult(code, "", data)
}

func (s *Result) Result() (err string) {
	return s.Message
}
