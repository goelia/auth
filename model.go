package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AuthModel abastract
type AuthModel struct {
	UserID uint `json:"user_id,omitempty"`
	IP     string
	gorm.Model
}

// User struct
type User struct {
	gorm.Model

	CreatedIP string    `json:"created_ip,omitempty"`
	SigninIP  string    `json:"signin_ip,omitempty"`
	SigninAt  time.Time `json:"signin_at,omitempty"`
	Username  string    `json:"username,omitempty" validate:"required,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Email     string    `json:"email,omitempty" sql:"unique_index:idx_email_mobile" validate:"email"`
	Mobile    string    `json:"mobile,omitempty" sql:"unique_index:idx_email_mobile"` //`validate:"mobile"`
	Name      string    `json:"name,omitempty"`
	Age       uint8     `json:"age,omitempty" validate:"gte=0,lte=130"`
	Birthday  time.Time `json:"birthday,omitempty"`

	Roles []Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

// TableName return User's table name
func (User) TableName() string {
	return "users"
}

// Role struct
type Role struct {
	gorm.Model

	Name        string       `json:"name,omitempty" sql:"type:varchar(100);unique_index"`
	Alias       string       `json:"alias,omitempty"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
}

// Permission struct
type Permission struct {
	gorm.Model

	Name  string `json:"name,omitempty" sql:"type:varchar(100);unique_index"`
	Alias string `json:"alias,omitempty"`
}

// AuthLocal struct
// ID is mobile or email
type AuthLocal struct {
	AuthModel
	Name     string `json:"name,omitempty" sql:"type:varchar(100);unique_index"` //email or mobile
	Password string `json:"password,-"`
	Salt     string `json:"-"`
	Code     string `json:"code,omitempty"` //校验码
	CodeAt   time.Time
	Expires  int `json:"expires,omitempty"`
}

// AuthCode 验证码记录
type AuthCode struct {
	gorm.Model
	Name    string `json:"name,omitempty" sql:"type:varchar(100);"` //email or mobile
	Code    string `json:"code,omitempty"`                          //校验码
	Content string `json:"content,omitempty"`                       //html or text
	Source  uint   //1:登录,2:绑定手机 3:绑定邮箱
	Expires int    `json:"expires,omitempty"`
}

// AuthAPI api登录
type AuthAPI struct {
	AuthModel
	APIKey    string `json:"key"`
	APISecret string `json:"secret"`
	Expires   uint   `json:"expires,omitempty"`
}

// OAuth 授权登录
type OAuth struct {
	AuthModel
	OAuthName        string `json:"oauth_name,omitempty" gorm:"column:oauth_name"`
	OAuthID          string `json:"oauth_id,omitempty" gorm:"column:oauth_id"`
	OAuthAccessToken string `json:"access_token,omitempty" gorm:"column:access_token"`
	Expires          uint   `json:"expires,omitempty" gorm:"column:expires"`
}

// TableName return oauth's table name
func (OAuth) TableName() string {
	return "oauth"
}
