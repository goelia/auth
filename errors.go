package auth

import "github.com/go-errors/errors"

var (
	// RequiredErr err
	RequiredErr = errors.Errorf("必填项不能为空")
	// ExpiresRefreshAuthCodeErr not expired
	ExpiresRefreshAuthCodeErr = errors.Errorf("请稍后重试")
	// ExpiresAuthCodeErr err
	ExpiresAuthCodeErr = errors.Errorf("验证码已过期")
)
