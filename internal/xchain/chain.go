package xchain

import "strings"

type Chain string

func (c Chain) Normalize() Chain {
	return Chain(strings.ToUpper(string(c)))
}
