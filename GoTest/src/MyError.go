package src

import (
	"error"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

//实现接口
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}
