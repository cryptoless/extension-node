package model

import "fmt"

// type MethodNotFoundError struct {
// 	Method string
// }

func (e *MethodNotFoundError) ErrorCode() int {
	return -1000
}
func (e *MethodNotFoundError) Error() string {
	return fmt.Sprintf("method %s does not exist/is not available", e.Method)
}

///
// type ParseError struct {
// 	Msg string
// }

func (e *ParseError) ErrorCode() int {
	return -1001
}
func (e *ParseError) Error() string {
	return e.Msg
}

///
// type InvalidRequestError struct {
// 	Msg string
// }

func (e *InvalidRequestError) ErrorCode() int {
	return -1002
}
func (e *InvalidRequestError) Error() string {
	return e.Msg
}

///
// type InvalidMessageError struct {
// 	Msg string
// }

func (e *InvalidMessageError) ErrorCode() int {
	return -1002
}
func (e *InvalidMessageError) Error() string {
	return e.Msg
}

///
// type InvalidParamsError struct {
// 	Msg string
// }

func (e *InvalidParamsError) ErrorCode() int {
	return -1002
}
func (e *InvalidParamsError) Error() string {
	return e.Msg
}

// type InternalError struct {
// 	Msg string
// }
func (e *InternalError) ErrorCode() int {
	return -1003
}
func (e *InternalError) Error() string {
	return e.Msg
}
