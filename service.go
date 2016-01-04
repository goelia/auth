package auth

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
)

// RefreshCode sending code
func (a *AuthLocal) RefreshCode(send func() error) error {
	if a.Name == "" {
		return errors.New(RequiredErr)
	}
	err := DB.Where("name = ?", a.Name).First(a).Error
	if err != nil && err != gorm.RecordNotFound {
		return err
	}
	a.Expires = ExpiresAuthCode
	code := strconv.Itoa(RandNum())
	tx := DB.Begin()
	if err == gorm.RecordNotFound {
		user := User{
			Email: a.Name,
		}
		tx.Create(&user)

		a.UserID = user.ID
		a.Code = code
		a.CodeAt = time.Now()
		if err = tx.Create(a).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if time.Now().Before(a.CodeAt.Add(time.Duration(ExpiresRefreshAuthCode) * time.Second)) {
			tx.Rollback()
			return errors.New(ExpiresRefreshAuthCodeErr)
		}
		a.Code = code
		a.CodeAt = time.Now()
		if err = tx.Save(a).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err = send(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Signin by local
func (a *AuthLocal) Signin() error {
	name := a.Name
	password := a.Password
	code := a.Code

	switch {
	case name != "" && password != "":
		h := sha256.New()
		io.WriteString(h, a.Password)
		a.Password = string(h.Sum(nil))
		fmt.Printf("% x", a.Password)
		if err := DB.Where("name=? and password=?", name, a.Password).First(a).Error; err != nil {
			return err
		}
	case name != "" && code != "":
		if err := DB.Where("name=? and code=?", name, code).First(a).Error; err != nil {
			return err
		}
		if a.UpdatedAt.Add(time.Second * time.Duration(a.Expires)).Before(time.Now()) { //判断校验码是否过期
			return errors.New(ExpiresAuthCodeErr)
		}
	default:
		return errors.New(RequiredErr)
	}
	tx := DB.Begin()
	if err := DB.Save(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	user := User{}
	tx.First(&user, a.UserID)
	user.SigninAt = time.Now()
	user.SigninIP = a.IP

	if err := DB.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
