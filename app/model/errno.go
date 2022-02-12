package model

type Error interface {
	ErrorCode() int
	Error() string
}
type MethodNotFoundError struct {
	Method string
}

///
type ParseError struct {
	Msg string
}

///
type InvalidRequestError struct {
	Msg string
}

///
type InvalidMessageError struct {
	Msg string
}

///
type InvalidParamsError struct {
	Msg string
}

///
type InternalError struct {
	Msg string
}
