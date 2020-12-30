/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 **/
package codeerror

var CodeError ErrorStruct



type ErrorWithCode struct {
	Code    int
	Message string
}

func (r ErrorWithCode) Error() string {
	return r.Message
}

func (r ErrorWithCode) GetCode() int{
	return r.Code
}



type ErrorStruct struct {
}
func  (e ErrorStruct)NewError(code int,msg string) ErrorWithCode {
	return ErrorWithCode{
		Code: code,
		Message: msg,
	}
}

func (e ErrorStruct) Assert(err error) ErrorWithCode {
	return err.(ErrorWithCode)
}

func init()  {
	CodeError = ErrorStruct{}
}
