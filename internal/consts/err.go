package consts

import "errors"

var (
	ErrUserNotFound = errors.New("用户不存在或用户密码错误")

	ErrTagNotFound = errors.New("标签未找到")
	ErrTagExist    = errors.New("标签已存在")
)
