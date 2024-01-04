package errno

import "fmt"

type Errno struct {
	Status  int
	Code    string
	Message string
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) SetMsg(format string, args ...interface{}) *Errno {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func DeCode(err error) (int, string, string) {
	if err == nil {
		return 200, "", "OK"
	}
	switch typed := err.(type) {
	case *Errno:
		return typed.Status, typed.Code, typed.Message
	default:
	}
	return InternalServerError.Status, InternalServerError.Code, InternalServerError.Message
}
