package utils

import "net/http"

// 定义自定义的错误类型
type CustomError struct {
	code     int    // 错误码
	msg      string // 错误提示
	httpCode int
}

// 实现 error 接口的 Error 方法
func (e *CustomError) Error() string {
	return e.msg
}

func (e *CustomError) Wrap(err error) *CustomError {
	if err == nil {
		return e
	}
	Msg := e.msg + "\n" + err.Error()
	if ce, ok := err.(*CustomError); ok {
		return getNewError(*e, ce.code, Msg)
	}

	return getNewError(*e, e.code, Msg)
}

func (e *CustomError) AddMsg(msg string) *CustomError {
	Msg := e.msg + "\n" + msg
	return getNewError(*e, e.code, Msg)
}

func getNewError(e CustomError, code int, Msg string) *CustomError {
	e.msg = Msg
	e.code = code
	return &e
}

func (e *CustomError) HttpReturn(writer http.ResponseWriter) {
	http.Error(writer,
		NewFormatResponseWithMap(e, nil), e.httpCode)
}

// 预定义错误
var (
	ErrGetUserInfoFailed             = &CustomError{code: 1001, msg: "获取用户信息/鉴权失败", httpCode: http.StatusPreconditionFailed}
	ErrTaskIDNotFound                = &CustomError{code: 1005, msg: "任务ID获取失败", httpCode: http.StatusPreconditionRequired}
	ErrRunningTaskNotFound           = &CustomError{code: 1006, msg: "查找不到指定在运行中的任务", httpCode: http.StatusExpectationFailed}
	ErrTaskNotFound                  = &CustomError{code: 1007, msg: "无法查找到对应的任务", httpCode: http.StatusExpectationFailed}
	ErrProjectIDNotFound             = &CustomError{code: 1008, msg: "项目ID获取失败", httpCode: http.StatusPreconditionRequired}
	ErrAuthenticationFailed          = &CustomError{code: 1009, msg: "鉴权失败", httpCode: http.StatusUnauthorized}
	ErrJsonMarshalerFailed           = &CustomError{code: 1010, msg: "结果Json转换失败", httpCode: http.StatusBadRequest}
	ErrRequestBodyDecodeFailed       = &CustomError{code: 1011, msg: "Post请求体解析失败", httpCode: http.StatusPaymentRequired}
	ErrJsonUnMarshalerFailed         = &CustomError{code: 1012, msg: "请求JSON转换", httpCode: http.StatusPaymentRequired}
	ErrDatasetNameDuplicate          = &CustomError{code: 1013, msg: "资源重复建立", httpCode: http.StatusConflict}
	ErrTransmitFailed                = &CustomError{code: 1014, msg: "发送请求时出错", httpCode: http.StatusFailedDependency}
	ErrRequestThirdServiceNotSuccess = &CustomError{code: 1015, msg: "第三方服务响应失败", httpCode: http.StatusFailedDependency}
	ErrReadResponseFailed            = &CustomError{code: 1016, msg: "读取响应时出错", httpCode: http.StatusBadRequest}
	ErrDecodeIntoMapFailed           = &CustomError{code: 1017, msg: "结果转换为Map时报错", httpCode: http.StatusBadRequest}
	ErrUpdateTaskInfoFailed          = &CustomError{code: 1018, msg: "更新任务信息失败", httpCode: http.StatusBadRequest}
	ErrProxyFailed                   = &CustomError{code: 1019, msg: "代理服务器连接失败", httpCode: http.StatusProxyAuthRequired}
)
