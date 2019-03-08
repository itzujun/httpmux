package error

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code int
	Msg  interface{}
	Err  error
}

func NewHttpError(code int, msg ...interface{}) *HttpError {
	e := &HttpError{
		Code: code,
		Msg:  http.StatusText(code),
	}
	if len(msg) > 0 {
		e.Msg = msg[0]
	}

	return e
}

func (e *HttpError) Error(err error) string {
	return fmt.Sprintf("code = %d,msg=%v", e.Code, e.Msg)
}

func (e* HttpError)SetInternal(err error)*HttpError{
	e.Err =err
	return e
}




