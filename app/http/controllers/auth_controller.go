package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

//Register注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

//DoRegister
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	//初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	//表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {

		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")

	} else {
		//验证成功
		_user.Create()

		if _user.ID > 0 {
			http.Redirect(w, r, "/", http.StatusFound)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}
}
