package auth

import (
	"errors"
	"goblog/app/models/user"
	"goblog/pkg/session"

	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}

	return ""
}

func User() user.User {
	uid := _getUID()

	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}

	return user.User{}
}

//尝试登录
func Attempt(email string, password string) error {

	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("内部错误，请稍后尝试")
		}

	}
	if !_user.ComparePassword(password) {

		return errors.New("账号不存在或密码错误")
	}

	session.Put("uid", _user.GetStringID())

	return nil
}

//login 登录指定用户
func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

// logout 退出用户
func Logout() {
	session.Forget("uid")
}

//检测是否登录
func Check() bool {
	return len(_getUID()) > 0
}
