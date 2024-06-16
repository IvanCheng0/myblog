package errno

import "fmt"

type Errno struct {
	Code    string
	Message string
	HTTP    int
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) SetMessage(format string, args ...interface{}) *Errno {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

// Decode 尝试从 err 中解析出业务错误码和错误信息.
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
