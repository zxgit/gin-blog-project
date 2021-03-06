/*
@Time : 2021/1/24 17:11
@Author : zhangxin
*/
package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code  int      `json:"code"`
	msg   string   `json:"msg"`
	trace []string `json:"trace"`
}

var codes = map[int]string{}

func NewErrow(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.trace
}
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.trace = []string{}
	for _, d := range details {
		newError.trace = append(newError.trace, d)
	}
	return &newError
}

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
