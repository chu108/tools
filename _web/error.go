package _web

import "errors"

var (
	ErrorNullDoc = errors.New("document是空的")
	ErrorNullSel = errors.New("selection是空的")
)
