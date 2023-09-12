package ecode

import (
	"fmt"
	"strconv"
)

var (
	_codes = map[int64]struct{}{} // register codes.
)

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int64) Code {
	if e <= 0 {
		panic("ecode must greater than zero")
	}
	return add(e)
}

func add(e int64) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Code(e)
}

type ErrCodes interface {
	Error() string
	Code() int64
	Message() string
}

type Code int64

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Code() int64 {
	return int64(e)
}

func (e Code) Message() string {
	return e.Error()
}
