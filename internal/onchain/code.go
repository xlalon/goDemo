package onchain

import "strings"

type Code string

func (c Code) Normalize() Code {
	return Code(strings.ToUpper(string(c)))
}
