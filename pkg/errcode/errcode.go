package errcode

import (
	"fmt"
	"net/http"
)

//Error 错误码结构体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details`
}

var codes = map[int]string{}

//NewError 新建一个错误码
func NewError(code int, msg string) *Error {
	_, ok := codes[code]
	if ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请换一个。", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

//Code 获取错误码
func (e *Error) Code() int {
	return e.code
}

//Msg 获取错误码的详细信息
func (e *Error) Msg() string {
	return e.msg
}

//Msgf 写入错误详细信息
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

//WithDetails 写入多个错误信息
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

//Details 获取当前错误信息
func (e *Error) Details() []string {
	return e.details
}

//StatusCode 获取http标准状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
