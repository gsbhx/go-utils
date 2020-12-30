package errors

import codeerror "github.com/gsbhx/tools/errors/error"
var Errors *AllErrors

type AllErrors struct {
	ErrorWithCode codeerror.ErrorStruct
}

func init()  {
	Errors=new(AllErrors)
}