package xerr

import (
	"fmt"
	"runtime"
)

// Xerr define refrence rest style
type Xerr struct {
	Code  int    `json:"code"`  // 200/500/401
	Msg   string `json:"msg"`   // message
	Cause string `json:"cause"` //golang code
}

func (e Xerr) Error() string {
	return fmt.Sprintf("Xerr-[Code:%v,Msg:%v,Cause:%v]", e.Code, e.Msg, e.Cause)
}

// New instance
func New(code int, msg string) (e Xerr) {

	//func Caller(skip int) (pc uintptr, file string, line int, ok bool) {
	_, fName, LN, _ := runtime.Caller(1)
	//e.Cause = fmt.Sprint(runtime.Caller(1))
	e.Cause = fmt.Sprintf("fName-%s,LN-%d", fName, LN)
	//runtime.Stack()
	e.Code = code
	e.Msg = msg

	return e
}
