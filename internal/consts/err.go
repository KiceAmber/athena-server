package consts

import "errors"

var (
	ErrTagNotFound = errors.New("标签未找到")
	ErrTagExist    = errors.New("标签已存在")
)
