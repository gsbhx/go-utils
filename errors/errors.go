package errors

import codeerror "tools/errors/error"
var Errors *AllErrors

type AllErrors struct {
	ErrorWithCode codeerror.ErrorStruct
}

func init()  {
	Errors=new(AllErrors)
}