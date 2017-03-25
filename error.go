package goutility

import "fmt"

const (
	UnspecifiedCode        = -1
	UnspecifiedCodeContext = "goutility"
)

type ErrorCode int

type ErrorType struct {
	message string

	code        ErrorCode
	codeContext string
}

func MakeError(message string) ErrorType {
	return MakeErrorWithCode(message, UnspecifiedCode, UnspecifiedCodeContext)
}

func MakeErrorWithCode(message string, code ErrorCode, codeContext string) ErrorType {
	return ErrorType{
		message:     message,
		code:        code,
		codeContext: codeContext,
	}
}

func (this ErrorType) Error() string {
	return fmt.Sprintf("Code: %d, Context: %s, %s", this.Code(), this.CodeContext(), this.Message())
}

func (this ErrorType) Message() string {
	return this.message
}

func (this ErrorType) Code() ErrorCode {
	return this.code
}

func (this ErrorType) CodeContext() string {
	return this.codeContext
}

func (this ErrorType) HasValidCode() bool {
	return this.Code() != UnspecifiedCode
}
