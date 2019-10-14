package errscan

import "errors"

var (
	WrongStatusCodeError = errors.New("wrong status code")
	EmptyResultError     = errors.New("nothing found")
	WrongMonthTypeError  = errors.New("month should be int type")
	WrongDayTypeError    = errors.New("day should be int type")
	EmptyDayValueError   = errors.New("day should be fill out")
	EmptyMonthValueError = errors.New("month should be fill out")
	BigMonthValueError   = errors.New("too big month")
	BigDayValueError     = errors.New("too big day")
)
