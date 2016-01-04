package auth

import "github.com/jinzhu/gorm"

var (
	// DB by gorm
	DB *gorm.DB
	// ExpiresAuthCode 校验码过期时间，单位：秒
	ExpiresAuthCode = 10 * 60 //10分钟
	// ExpiresRefreshAuthCode 校验码刷新时间间隔
	ExpiresRefreshAuthCode = 2 * 60 //2分钟
)
