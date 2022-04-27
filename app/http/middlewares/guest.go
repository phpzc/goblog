package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

//只允许未登录用户访问
func Guest(next HttpHandlerFunc) HttpHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if auth.Check() {
			flash.Warning("登录用户无法访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next(w, r)
	}
}