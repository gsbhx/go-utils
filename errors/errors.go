package errors

import codeerror "git.hsuanyuen.cn/tools/errors/error"
var Errors *AllErrors

type AllErrors struct {
	ErrorWithCode codeerror.ErrorStruct
}

func init()  {
	Errors=new(AllErrors)
}