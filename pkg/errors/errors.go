package errors

import (
	"moocss.com/gaea/pkg/twirp"

	"github.com/pkg/errors"
)

var (
	// NotLoginError 错误未登录
	NotLoginError = twirp.NewError(twirp.Unauthenticated, "must login")

	// PermissionDeniedError 权限不够
	PermissionDeniedError = twirp.NewError(twirp.PermissionDenied, "permission denied")

	// token过期
	TokenExpired error = twirp.NewError(twirp.DeadlineExceeded, "Token is expired")

	// token未激活
	TokenNotValidYet error = twirp.NewError(twirp.InvalidArgument, "Token not active yet")

	// 不是token
	TokenMalformed error = twirp.NewError(twirp.InvalidArgument, "That's not even a token")

	// 无法处理token
	TokenInvalid error = twirp.NewError(twirp.InvalidArgument, "Couldn't handle this token:")

	// token生成错误
	TokenFailure error = twirp.NewError(twirp.InvalidArgument, "Token generate failure")
)

// Wrap 包装错误信息，附加调用栈
// 第二个参数只能是 string，也可以不传，大部分情况不用传
func Wrap(err error, args ...interface{}) error {
	if len(args) >= 1 {
		if msg, ok := args[0].(string); ok {
			return errors.Wrap(err, msg)
		}
	}

	return errors.Wrap(err, "")
}

// Cause 获取原始错误对象
func Cause(err error) error {
	return errors.Cause(err)
}

// Errorf 创建新错误
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// InvalidArgumentError 参数错误，400
func InvalidArgumentError(argument string, validationMsg string) error {
	return twirp.InvalidArgumentError(argument, validationMsg)
}

type codeError struct {
	code int32
	err  string
}

func (c codeError) Error() string {
	return c.err
}

// CodeError 新建业务错误，附带错误码
func CodeError(code int32, err string) error {
	return codeError{code: code, err: err}
}

// Code 提取错误码，codeError 返回 code 和 true，其他返回 0 和 false
func Code(err error) (int32, bool) {
	err = errors.Cause(err)

	if err == nil {
		return 0, false
	}

	if ce, ok := err.(codeError); ok {
		return ce.code, true
	}

	return 0, false
}
